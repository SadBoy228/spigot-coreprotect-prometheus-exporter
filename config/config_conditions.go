package config

import (
    "strings"

    "github.com/BurntSushi/toml"
)

type validCondition interface {
    Check(*toml.MetaData, *ApplicationConfig)  error
}

// Condition that checks defined metadata
type validConditionDefinedMetadata struct {
    metadataString string
}

func (dm validConditionDefinedMetadata) Check(metadata *toml.MetaData, conf *ApplicationConfig) error {
    if !metadata.IsDefined(strings.Split(dm.metadataString, ".")...) {
        return &ConfigOptionIsNotDefined{dm.metadataString}
    }

    return nil
}

// Condition type that wraps arbitrary function
type validConditionFunc func(*toml.MetaData, *ApplicationConfig) error

func (f validConditionFunc) Check(m *toml.MetaData, c *ApplicationConfig) error {
    return f(m, c)
}

// Private function that checks database options
func checkDatabaseOptions(metadata *toml.MetaData, conf *ApplicationConfig) error {

    var optionsToCheck []string

    switch conf.DB.DatabaseType {
    case "sqlite":
        optionsToCheck = []string{
            "db.sqlite_file_path",
        }
    case "mysql":
        optionsToCheck = []string{
            "db.mysql_hostname",
            "db.mysql_port",
            "db.mysql_database",
            "db.mysql_usename",
            "db.mysql_password",
        }
    default:
        return ErrInvalidDatabaseType
    }

    for _, optionToCheck := range optionsToCheck {
        if !metadata.IsDefined(strings.Split(optionToCheck, ".")...) {
            return &ConfigOptionIsNotDefined{optionToCheck}
        }
    }

    return nil
}

// Private function that checks http SSL options
func checkSSLOptions(metadata *toml.MetaData, conf *ApplicationConfig) error {

    if conf.HTTP.UseSSL {
        for _, optionToCheck := range []string{"http.cert_file", "http.key_file"} {
            if !metadata.IsDefined(strings.Split(optionToCheck, ".")...) {
                return &ConfigOptionIsNotDefined{optionToCheck}
            }
        }
    }

    return nil
}
