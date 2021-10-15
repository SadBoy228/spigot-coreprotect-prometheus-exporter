package main

import (
    "log"
    "context"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"
)

func main() {

    _, cancel := context.WithCancel(context.Background())
    cfg, err := config.GetConfiguration()

    if err != nil {
        log.Println(err)
        return
    }

    _, err = appLog.CreateLogger(cancel, "Application", appLog.LoggerConfig{
        OutputLogFile: cfg.OutputLogFile,
        ErrorLogFile: cfg.ErrorLogFile,
        EnableDebugLog: cfg.EnableDebugLog,
    })

    if err != nil {
        log.Println(err)
        return
    }
}
