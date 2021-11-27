package main

import (
    "log"
    "context"
    "os"
    "os/signal"
    "syscall"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/db"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/service"
)

func main() {

    // Init configuration
    cfg, err := config.InitConfiguration()

    if err != nil {
        log.Println(err)
        return
    }

    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh,
        os.Interrupt,
        syscall.SIGHUP,
        syscall.SIGTERM,
    )

    appCtx, appCancel := context.WithCancel(context.Background())
    mainLogger, err := appLog.CreateLogger(appCancel, cfg, "Application")

    if err != nil {
        log.Println(err)
        return
    }

    // Initialize database
    db, err := db.InitDatabaseConnection(cfg)

    if err != nil {
        mainLogger.Fatal("Error when initializing database: ", err)
        return
    }

    // Start goroutines
    closeNotifier := service.RunServices(appCtx, appCancel, cfg, db)

    go func() {
        <-signalCh
        mainLogger.Info("Got signal, stopping...")
        appCancel()
    }()

    closeNotifier.Wait()
}
