//+build mage

package main

const (
    // Path to directory, where build artifacts will be stored after build
    buildPath = "build/"

    // Name of generated executable file
    executableName = "cp-prometheus-exporter"

    // Path to directory, where to install executable file
    installPrefix = "/usr/local/bin/"

    // Path to directory, where to install configuration file
    configurationFileLocation = "/etc/cp-prometheus-exporter/"

    // Enable Systemd unit file installation
    // If you are on BSD or OSX or doesn't have Systemd, change this to false
    installSystemdService = true
)
