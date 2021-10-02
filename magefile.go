//+build mage

package main

import (
    "strings"
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
    return sh.RunV("go", "build", "-o", buildPath + executableName)
}

func Install() error {
    if _, err := os.Stat(buildPath + executableName); os.IsNotExist(err) {
        mg.Deps(Build)
    } else {
        return err
    }

    return sh.Copy(installPrefix + executableName, buildPath + executableName)
}
