package config

import (
    "github.com/BurntSushi/toml"
)

type ApplicationConfig struct {
    UpdateIntervalSec       uint                `toml:"update_interval"`
    EnableDebugLog          bool                `toml:"enable_debug_log"`
    OutputLogFile           string              `toml:"output_log_file"`
    ErrorLogFile            string              `toml:"error_log_file"`
    DB                      DatabaseConfig      `toml:"db"`
    HTTP                    HttpServerConfig    `toml:"http"`

    parseMetadata           toml.MetaData
    validConfigConditions   []validCondition
}

func (a *ApplicationConfig) Parse(data string) (errSlice []error) {

    var err error
    a.parseMetadata, err = toml.Decode(data, a)

    if err != nil {
        errSlice = append(errSlice, err)
        return
    }

    errSlice = a.checkConditions()
    return
}

func (a *ApplicationConfig) GetMetadata() *toml.MetaData {
    return &a.parseMetadata
}

func (a *ApplicationConfig) checkConditions() []error {

    resultErrors := make([]error, 0, len(a.validConfigConditions))

    for _, condition := range a.validConfigConditions {
        if err := condition.Check(&a.parseMetadata, a); err != nil {
            resultErrors = append(resultErrors, err)
        }
    }

    return resultErrors
}

type DatabaseConfig struct {
    DatabaseType        string `toml:"database_type"`
    TablePrefix         string `toml:"table_prefix"`
    SqliteDatabasePath  string `toml:"sqlite_file_path"`
    MySQLHostname       string `toml:"mysql_hostname"`
    MySQLPort           string `toml:"mysql_port"`
    MySQLDatabaseName   string `toml:"mysql_database"`
    MySQLUsername       string `toml:"mysql_username"`
    MySQLPassword       string `toml:"mysql_password"`
}

type HttpServerConfig struct {
    ListenAddr  string  `toml:"listen_addr"`
    ListenPort  uint64  `toml:"listen_port"`
    UseSSL      bool    `toml:"use_ssl"`
    CertPath    string  `toml:"cert_file"`
    KeyPath     string  `toml:"key_file"`
}
