package db

import (
	"fmt"

	"github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
    db *gorm.DB
)

func GetDatabaseConnection() (*gorm.DB, error) {
    var err error

    if db == nil {
        cfg, _ := config.GetConfiguration()

        if cfg.DB.DatabaseType == "sqlite" {
            db, err = gorm.Open(sqlite.Open(cfg.DB.SqliteDatabasePath), &gorm.Config{})
        } else if cfg.DB.DatabaseType == "mysql" {
            dsn := fmt.Sprintf(
                "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
                cfg.DB.MySQLUsername,
                cfg.DB.MySQLPassword,
                cfg.DB.MySQLHostname,
                cfg.DB.MySQLPort,
                cfg.DB.MySQLDatabaseName,
            )

            db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        } else {
            panic("Unknown database type has defined")
        }
    }

    return db, err
}
