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

// Build script configuration
const (
    buildPath = "build/"
    executableName = "cp-prometheus-exporter"
    installPrefix = "/usr/local/bin/"

    configurationFileLocation = "/etc/cp-prometheus-exporter/"

    installSystemdService = true // If you are on BSD or OSX, change this to false, or doesn't have Systemd, change this to false
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
)

func Build() error {
    return sh.RunV("go", "build", "-o", buildPath + executableName)
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

    if err := sh.Copy("/etc/systemd/system/cp-prometheus-exporter.service", buildPath + "cp-prometheus-exporter.service"); err != nil {
        return err
    }

    return sh.RunV("systemctl", "daemon-reload")
}
