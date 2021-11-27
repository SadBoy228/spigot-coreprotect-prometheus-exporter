package db_agregator

import (
    "context"
    "sync"
    "time"
    "log"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"

    "gorm.io/gorm"
)

type DBAgregatorService struct {}

func (d *DBAgregatorService) StartService(
    appCtx context.Context,
    appCancel context.CancelFunc,
    closeNotifier *sync.WaitGroup,
    cfg *config.ApplicationConfig,
    db *gorm.DB,
) {
    agregatorLogger, err := appLog.CreateLogger(appCancel, cfg, "Agregator")

    if err != nil {
        log.Println(err)
        appCancel()
        return
    }

    ticker := time.NewTicker(time.Duration(cfg.UpdateIntervalSec) * time.Second)

    for {
        select {
        case <-ticker.C:
            performAgregation(appCtx, appCancel, agregatorLogger)
        case <-appCtx.Done():
            agregatorLogger.Info("Stopping db agregation service...")
            ticker.Stop()
            return
        }
    }
}

func performAgregation(
    appCtx context.Context,
    appCancel context.CancelFunc,
    logger *appLog.Logger,
) {
}
