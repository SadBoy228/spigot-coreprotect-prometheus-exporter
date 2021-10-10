package db_connection

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeConnection() *gorm.DB {
	cfg, _ := config.GetConfiguration()

	var db *gorm.DB
	var err error

	if cfg.DB.DatabaseType == "sqlite" {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	} else {
		dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB.MySQLHostname, cfg.DB.MySQLPassword, cfg.DB.MySQLPort, cfg.DB.MySQLDatabaseName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("error: ", err)
	}

	return db
}
