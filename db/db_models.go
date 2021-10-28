package db

import (
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/db/modeltypes"

    "github.com/google/uuid"
)

type ChatMessage struct{
    RowID       int                 `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime `gorm:"column:time;index:time;index:user;index:wid"`
    UserID      int                 `gorm:"column:user;index:user"`
    WorldID     int                 `gorm:"column:wid;index:wid"`

    X           int                 `gorm:"column:x;index:wid"`
    Y           int                 `gorm:"column:y"`
    Z           int                 `gorm:"column:z;index:wid"`

    Message     string              `gorm:"column:message;size:16000"`
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

type Session struct{
    RowID       int                             `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime             `gorm:"column:time;index:time;index:user;index:wid;index:action"`
    UserID      int                             `gorm:"column:user;index:user"`
    WorldID     int                             `gorm:"column:wid;index:wid"`

    X           int                             `gorm:"column:x;index:wid"`
    Y           int                             `gorm:"column:y"`
    Z           int                             `gorm:"column:z;index:wid"`

    Action      modeltypes.SessionActionType    `gorm:"column:action;index:action"`
}

func (s Session) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "session"
}

type Command struct {
    RowID       int                 `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime `gorm:"column:time;index:time;index:user;index:wid"`
    UserID      int                 `gorm:"column:user;index:user"`
    WorldID     int                 `gorm:"column:wid;index:wid"`

    X           int                 `gorm:"column:x;index:wid"`
    Y           int                 `gorm:"column:y"`
    Z           int                 `gorm:"column:z;index:wid"`

    Message     string              `gorm:"column:message;size:16000"`
}

func (c Command) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "command"
}

type Block struct {
    RowID       int64                       `gorm:"column:rowid;primaryKey;autoIncrement"`
    Timestamp   modeltypes.Unixtime         `gorm:"column:time;index:type;index:user;index:wid"`
    UserID      int                         `gorm:"column:user;index:user"`
    WorldID     int                         `gorm:"column:wid;index:wid"`

    X           int                         `gorm:"column:x;index:wid"`
    Y           int                         `gorm:"column:y"`
    Z           int                         `gorm:"column:z;index:wid"`

    Type        int                         `gorm:"column:type;index:type"`
    Data        int                         `gorm:"column:data"`

    Metadata    []byte                      `gorm:"column:meta;type:mediumblob"`

    Action      modeltypes.BlockActionType  `gorm:"column:action"`
    RolledBack  modeltypes.BoolTinyInt      `gorm:"column:rolled_back;type:tinyint(1)"`
}

func (b Block) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "block"
}

type BlockMetadataMap struct {
    RowID   int     `gorm:"column:rowid;primaryKey;autoIncrement"`
    ID      int     `gorm:"column:id;index:id"`
    Data    string  `gorm:"column:data;size:255"`
}

func (b BlockMetadataMap) TableName() string {
    cfg, _ := config.GetConfiguration()
    return cfg.DB.TablePrefix + "blockdata_map"
}
