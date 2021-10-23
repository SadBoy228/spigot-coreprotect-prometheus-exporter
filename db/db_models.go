package db

import (
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/db/model_types"

    "github.com/google/uuid"
)

type ChatMessage struct{
    RowID       int                 `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime `gorm:"column:time;index:time;index:user;index:wid"`
    UserID      int                 `gorm:"column:user;index:user"`
    WorldID     int                 `gorm:"column:wid;index:wid"`

    X   int `gorm:"column:x;index:wid"`
    Y   int `gorm:"column:y"`
    Z   int `gorm:"column:z;index:wid"`

    Message string  `gorm:"column:message;size:16000;default:NULL"`
}

func (c ChatMessage) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "chat"
}

type World struct{
    RowID   int     `gorm:"column:rowid;primaryKey;autoIncrement"`
    ID      int     `gorm:"column:id;index:id"`
    World   string  `gorm:"column:world;size:255"`
}

func (w World) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "world"
}

type User struct{
    RowID       int                 `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime `gorm:"column:time"`
    UserID      int                 `gorm:"column:user;index:user"`
    UserUUID    uuid.UUID           `gorm:"column:uuid;index:uuid"`
}

func (u User) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "user"
}
