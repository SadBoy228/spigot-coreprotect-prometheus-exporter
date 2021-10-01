# Spigot CoreProtect Prometheus Exporter
[![License](https://img.shields.io/github/license/k0tletka/spigot-coreprotect-prometheus-exporter?&logo=github)](LICENSE)

This service exports minecraft metric data to Prometheus system monitoring in appropriate format, using collected data by CoreProtect spigot plugin.
Required version of Golang: `>=1.15`

Steps to build and install:
```
git clone https://github.com/k0tletka/spigot-coreprotect-prometheus-exporter
cd ./spigot-coreprotect-prometheus-exporter
go mod tidy
$(go env GOPATH)/bin/mage build
$(go env GOPATH)/bin/mage install
```
