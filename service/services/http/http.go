package http

import (
    "fmt"
    "log"
    "context"
    "sync"
    "net/http"
    "time"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/metrics"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"

    "github.com/gorilla/mux"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "gorm.io/gorm"
)

type HTTPService struct {
    serverErrorChannel chan error
    router *mux.Router
    httpserver *http.Server
}

func (h *HTTPService) initService(cfg *config.ApplicationConfig) {
    metricRegisterer := metrics.GetMetricRegistry()

    h.router = mux.NewRouter()
    h.router.Handle("/mertics", promhttp.HandlerFor(metricRegisterer, promhttp.HandlerOpts{}))

    h.httpserver = &http.Server{
        Addr: fmt.Sprintf("%s:%d", cfg.HTTP.ListenAddr, cfg.HTTP.ListenPort),
        Handler: h.router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        ReadHeaderTimeout: 10 * time.Second,
    }

    h.serverErrorChannel = make(chan error, 1)
}

func (h *HTTPService) StartService(
    appCtx context.Context,
    appCancel context.CancelFunc,
    closeNotifier *sync.WaitGroup,
    cfg *config.ApplicationConfig,
    db *gorm.DB,
) {
    if (h == nil) {
        panic("Method runs with nil receiver")
    }

    h.initService(cfg)

    defer closeNotifier.Done()
    defer close(h.serverErrorChannel)

    httpLogger, err := appLog.CreateLogger(appCancel, cfg, "Http")

    if err != nil {
        log.Println(err)
        appCancel()
        return
    }

    go func() {
        var e error

        if cfg.HTTP.UseSSL {
            e = h.httpserver.ListenAndServeTLS(cfg.HTTP.CertPath, cfg.HTTP.KeyPath)
        } else {
            e = h.httpserver.ListenAndServe()
        }

        h.serverErrorChannel <- e
    }()

    select {
    case <-appCtx.Done():
        httpLogger.Info("Stopping HTTP server...")
        h.httpserver.Shutdown(appCtx)

        <-h.serverErrorChannel
    case e := <-h.serverErrorChannel:
        httpLogger.Fatal("Error occured while serving http requests: ", e)
    }
}
