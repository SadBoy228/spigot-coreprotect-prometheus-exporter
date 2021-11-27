package log

import (
    "log"
    "os"
    "fmt"
    "context"
    "io"
    "runtime"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
)

const (
    infoPrefix  = "[INFO]:"
    warnPrefix  = "[WARN]:"
    errorPrefix = "[ERROR]:"
    fatalPrefix = "[FATAL]:"
    debugPrefix = "[DEBUG]:"
)

type Logger struct {
    outputFile      *os.File
    errorFile       *os.File
    outputLogger    *log.Logger
    errorLogger     *log.Logger

    cancel context.CancelFunc
    config *config.ApplicationConfig
}

func CreateLogger(cancel context.CancelFunc, cfg *config.ApplicationConfig, name string) (*Logger, error) {
    resultLogger := Logger{}
    var err error

    if len(cfg.OutputLogFile) != 0 {
        if resultLogger.outputFile, err = openLoggerFile(cfg.OutputLogFile); err != nil {
            return nil, err
        }

        resultLogger.outputLogger = registerNewLogger(
            name,
            io.MultiWriter(os.Stdout, resultLogger.outputFile),
        )
    }

    if len(cfg.ErrorLogFile) != 0 {
        if resultLogger.errorFile, err = openLoggerFile(cfg.ErrorLogFile); err != nil {
            return nil, err
        }

        resultLogger.errorLogger = registerNewLogger(
            name,
            io.MultiWriter(os.Stdout, resultLogger.errorFile),
        )
    }

    if resultLogger.outputLogger == nil {
        resultLogger.outputLogger = registerNewLogger(name, os.Stdout)
    }

    if resultLogger.errorLogger == nil {
        resultLogger.errorLogger = registerNewLogger(name, os.Stdout)
    }

    resultLogger.cancel = cancel
    resultLogger.config = cfg
    return &resultLogger, nil
}

func (l *Logger) Info(args ...interface{}) {
    l.logArgs(l.outputLogger, infoPrefix, args)
}

func (l *Logger) Warn(args ...interface{}) {
    l.logArgs(l.outputLogger, warnPrefix, args)
}

func (l *Logger) Error(args ...interface{}) {
    l.logArgs(l.errorLogger, errorPrefix, args)
}

func (l *Logger) Fatal(args ...interface{}) {
    l.logArgs(l.errorLogger, fatalPrefix, args)

    // Additionaly, close context
    l.cancel()
}

func (l *Logger) Debug(args ...interface{}) {
    if l.config.EnableDebugLog {
        l.logArgs(l.outputLogger, warnPrefix, args)
    }
}

func (l *Logger) logArgs(logger *log.Logger, logPrefix string, args []interface{}) {
    printArgs := []interface{}{logPrefix}
    printArgs = append(printArgs, args...)

    if logger != nil {
        logger.Println(printArgs...)
    }
}

func openLoggerFile(filename string) (*os.File, error) {
    fd, err := os.OpenFile(
        filename,
        os.O_WRONLY | os.O_CREATE | os.O_APPEND,
        0755,
    )

    if err != nil {
        return nil, err
    }

    runtime.SetFinalizer(fd, func(f interface{}) {
        file := f.(*os.File)
        file.Sync()
        file.Close()
    })

    return fd, nil
}

func registerNewLogger(name string, writer io.Writer) *log.Logger {
    return log.New(
        writer,
        fmt.Sprintf("<%s> ", name),
        log.Lmsgprefix | log.LstdFlags,
    )
}
