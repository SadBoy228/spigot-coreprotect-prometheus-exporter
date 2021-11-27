package service

import (
    "context"
    "sync"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/service/services/http"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/service/services/db_agregator"

    "gorm.io/gorm"
)

var (
    ApplicationServices = [...]ApplicationService{
        &http.HTTPService{},
        &db_agregator.DBAgregatorService{},
    }
)

// Interface of application service, that performs certain actions
// during all application runtime
type ApplicationService interface {
    // Function, that starts service. Can be only ran once. Arguments listed below:
    // appCtx - Application context, if appCtx.Done() is true,
    //     all ApplicationService instances must stop their work.
    //
    // appCancel - Function, that sets "Done" state to appCtx, notifying,
    //     that all other ApplicationService instance must stop their work.
    //
    // closeNotifier - WaitGroup object, used to notify main application code, that
    //     application can be closed.
    //
    // cfg - Config object.
    // db - Database object.
    StartService(
        appCtx context.Context,
        appCancel context.CancelFunc,
        closeNotifier *sync.WaitGroup,
        cfg *config.ApplicationConfig,
        db *gorm.DB,
    )
}

func RunServices(
    appCtx context.Context,
    appCancel context.CancelFunc,
    cfg *config.ApplicationConfig,
    db *gorm.DB,
) *sync.WaitGroup {
    closeNotifier := &sync.WaitGroup{}
    closeNotifier.Add(len(ApplicationServices))

    for _, as := range ApplicationServices {
        go as.StartService(appCtx, appCancel, closeNotifier, cfg, db)
    }

    return closeNotifier
}
