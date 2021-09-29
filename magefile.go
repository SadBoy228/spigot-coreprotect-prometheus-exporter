//+build mage

package main

import (
    "os"

    "github.com/magefile/mage/mg"
    "github.com/magefile/mage/sh"
)

const (
    buildPath = "build/"
    executableName = "cp-prometheus-exporter"
    installPrefix = "/usr/local/bin/"
)

func Build() error {
    if err := sh.RunV("go", "build", "-o", buildPath + executableName); err != nil {
        return err
    }

    return nil
}

func Install() error {
    if _, err := os.Stat(buildPath + executableName); os.IsNotExist(err) {
        mg.Deps(Build)
    } else {
        return err
    }

    if err := sh.Copy(installPrefix + executableName, buildPath + executableName); err != nil {
        return err
    }

    return nil
}
