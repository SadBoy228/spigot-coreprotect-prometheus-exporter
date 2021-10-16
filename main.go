package main

import (
    "log"
    "context"
    "os"
    "os/signal"
    "sync"
    "syscall"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/http"
)

func main() {

    cfg, err := config.GetConfiguration()

    if err != nil {
        log.Println(err)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    wg := sync.WaitGroup{}

    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh,
        os.Interrupt,
        syscall.SIGHUP,
        syscall.SIGTERM,
        syscall.SIGKILL,
    )

    mainLogger, err := appLog.CreateLogger(cancel, "Application", appLog.LoggerConfig{
        OutputLogFile: cfg.OutputLogFile,
        ErrorLogFile: cfg.ErrorLogFile,
        EnableDebugLog: cfg.EnableDebugLog,
    })

    if err != nil {
        log.Println(err)
        return
    }

    // Start goroutines
    wg.Add(1)
    go http.StartHTTPServer(ctx, cancel, &wg)

    go func() {
        <-signalCh
        mainLogger.Info("Got signal, stopping...")
        cancel()
    }()

    wg.Wait()
}
