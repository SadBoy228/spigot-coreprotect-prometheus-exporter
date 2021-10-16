package db

import (
	"fmt"
	"log"

	"github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeConnection() *gorm.DB {
	cfg, _ := config.GetConfiguration()

	var db *gorm.DB
	var err error

	if cfg.DB.DatabaseType == "sqlite" {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

	if err != nil {
		log.Fatalln("error: ", err)
	}

	return db
}
