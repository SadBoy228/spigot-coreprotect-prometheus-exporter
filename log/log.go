package log

import (
    "log"
    "os"
    "fmt"
    "context"
)

const (
    infoPrefix  = "[INFO]:"
    warnPrefix  = "[WARN]:"
    errorPrefix = "[ERROR]:"
    fatalPrefix = "[FATAL]:"
    debugPrefix = "[DEBUG]:"
)

type LoggerConfig struct {
    OutputLogFile string
    ErrorLogFile string
    EnableDebugLog bool
}

type Logger struct {
    outputLogger *log.Logger
    errorLogger *log.Logger
    stdoutLogger *log.Logger

    cancel context.CancelFunc
    config *LoggerConfig
}

func CreateLogger(cancel context.CancelFunc, name string, conf LoggerConfig) (*Logger, error) {
    resultLogger := Logger{}

    if len(conf.OutputLogFile) != 0 {
        if outputLogger, err := setupInternalFileLogger(name, conf.OutputLogFile); err != nil {
            return nil, err
        } else {
            resultLogger.outputLogger = outputLogger
        }
    }

    if len(conf.ErrorLogFile) != 0 {
        if errorLogger, err := setupInternalFileLogger(name, conf.ErrorLogFile); err != nil {
            return nil, err
        } else {
            resultLogger.errorLogger = errorLogger
        }
    }

    resultLogger.stdoutLogger = log.New(
        os.Stdout,
        fmt.Sprintf("<%s> ", name),
        log.Lmsgprefix | log.LstdFlags,
    )

    resultLogger.cancel = cancel
    resultLogger.config = &conf
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

func (l *Logger) logArgs(logger *log.Logger, logPrefix string, args ...interface{}) {
    printArgs := []interface{}{logPrefix}
    printArgs = append(printArgs, args...)

    if logger != nil {
        logger.Println(printArgs...)
    }

    l.stdoutLogger.Println(printArgs...)
}

func setupInternalFileLogger(name string, filename string) (*log.Logger, error) {
    logFd, err := os.OpenFile(
        filename,
        os.O_WRONLY | os.O_CREATE | os.O_APPEND,
        0755,
    )

    if err != nil {
        return nil, err
    }

    return log.New(
        logFd,
        fmt.Sprintf("<%s> ", name),
        log.Lmsgprefix | log.LstdFlags,
    ), nil
}
