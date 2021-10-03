package config

import (
    "log"
    "os"
)

var (
    conf *ApplicationConfig
)

func GetConfiguration() (*ApplicationConfig, error) {

    if conf == nil {
        // Parse configuration that not defined
        appconf := &ApplicationConfig{
            validConfigConditions: []validCondition{
                validConditionDefinedMetadata{"update_interval"},
                validConditionDefinedMetadata{"db.database_type"},
                validConditionDefinedMetadata{"http.listen_port"},
                validConditionDefinedMetadata{"http.use_ssl"},
                validConditionFunc(checkDatabaseOptions),
                validConditionFunc(checkSSLOptions),
            },
        }

        conffile, ok := os.LookupEnv("CONFFILE")

        if !ok {
            conffile = "config.toml"
        }

        fileData, err := os.ReadFile(conffile)

        if err != nil {
            return nil, err
        }

        errs := appconf.Parse(string(fileData))

        if len(errs) != 0 {
            for _, err := range errs {
                log.Println(err)
            }

            return nil, ErrConfigParseError
        }

        conf = appconf
    }

    return conf, nil
}
