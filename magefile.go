//+build mage

package main

import (
    "text/template"
    "bytes"
    "os"
    "path/filepath"

    "github.com/magefile/mage/mg"
    "github.com/magefile/mage/sh"
)

type ServiceTemplateData struct {
    ConfigurationFileLocation string
    InstallPrefix string
    ExecutableName string
}

var (
    serviceTmplData = ServiceTemplateData{
        configurationFileLocation,
        installPrefix,
        executableName,
    }

    filesToInstall = map[string]string{
        installPrefix + executableName: buildPath + executableName,
        configurationFileLocation + "config.toml": "config.toml",
    }

    buildEnv = map[string]string{
        "CGO_ENABLED": "0",
    }
)

func Build() error {
    return sh.RunWithV(buildEnv, "go", "build", "-o", buildPath + executableName)
}

func Install() error {
    if _, err := os.Stat(buildPath + executableName); os.IsNotExist(err) {
        mg.Deps(Build)
    } else {
        return err
    }

    if installSystemdService {
        mg.Deps(InstallSystemdService)
    }

    for dst, src := range filesToInstall {
        dstPath := filepath.Dir(dst)

        if _, err := os.Stat(dstPath); os.IsNotExist(err) {
            if err = os.MkdirAll(dstPath, 0755); err != nil {
                return err
            }
        } else if err != nil {
            return err
        }

        if err := sh.Copy(dst, src); err != nil {
            return err
        }
    }

    return nil
}

func GenServiceTemplate() error {
    // Execute Systemd service template
    tmpl := template.Must(template.ParseFiles("cp-prometheus-exporter.service.template"))

    buffer := &bytes.Buffer{}
    if err := tmpl.Execute(buffer, serviceTmplData); err != nil {
        return err
    }

    return os.WriteFile(buildPath + "cp-prometheus-exporter.service", buffer.Bytes(), 0664)
}

func Clean() error {
    return os.RemoveAll(buildPath)
}

func InstallSystemdService() error {
    mg.Deps(GenServiceTemplate)

    if err := sh.Copy("/usr/lib/systemd/system/cp-prometheus-exporter.service", buildPath + "cp-prometheus-exporter.service"); err != nil {
        return err
    }

    return sh.RunV("systemctl", "daemon-reload")
}
