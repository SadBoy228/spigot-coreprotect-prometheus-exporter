package http

import (
    "fmt"
    "log"
    "io"
    "context"
    "sync"
    "net/http"
    "time"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    appLog "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/log"

    "github.com/gorilla/mux"
)

func StartHTTPServer(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
    defer wg.Done()

    cfg, _ := config.GetConfiguration()
    httpLogger, err := appLog.CreateLogger(cancel, "Http", appLog.LoggerConfig{
        OutputLogFile: cfg.OutputLogFile,
        ErrorLogFile: cfg.ErrorLogFile,
        EnableDebugLog: cfg.EnableDebugLog,
    })

    if err != nil {
        log.Println(err)
        return
    }

    serverErrorChan := make(chan error, 1)
    defer close(serverErrorChan)

    router := mux.NewRouter()
    router.HandleFunc("/", sampleHandler).PathPrefix("/")

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

func sampleHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Sample message")
}
