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
)

func StartHTTPServer(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
    defer wg.Done()

    cfg, _ := config.GetConfiguration()
    httpLogger, err := appLog.CreateLogger(cancel, "Http")

    if err != nil {
        log.Println(err)
        cancel()
        return
    }

    metricRegisterer, err := metrics.GetMetricRegistry()

    if err != nil {
        httpLogger.Fatal("Error when registering prometheus metrics: ", err)
        return
    }

    serverErrorChan := make(chan error, 1)
    defer close(serverErrorChan)

    router := mux.NewRouter()
    router.Handle("/mertics", promhttp.HandlerFor(metricRegisterer, promhttp.HandlerOpts{}))

    server := &http.Server{
        Addr: fmt.Sprintf("%s:%d", cfg.HTTP.ListenAddr, cfg.HTTP.ListenPort),
        Handler: router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        ReadHeaderTimeout: 10 * time.Second,
    }

    go func() {
        var e error

        if cfg.HTTP.UseSSL {
            e = server.ListenAndServeTLS(cfg.HTTP.CertPath, cfg.HTTP.KeyPath)
        } else {
            e = server.ListenAndServe()
        }

        serverErrorChan <- e
    }()

    select {
    case <-ctx.Done():
        httpLogger.Info("Stopping HTTP server...")
        server.Shutdown(ctx)
        cancel()

        <-serverErrorChan
    case e := <-serverErrorChan:
        httpLogger.Fatal("Error occured while serving http requests: ", e)
    }
}
