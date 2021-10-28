package modeltypes

import (
    "fmt"
    "database/sql/driver"
)

const (
    BlockBreak BlockActionType = iota
    BlockPlace
    BlockClick
    BlockKill
)

type BlockActionType int

func (b BlockActionType) Value() (driver.Value, error) {
    return int(b), nil
}

func (b *BlockActionType) Scan(value interface{}) error {
    v, ok := value.(int)

    if ok {
        *b = BlockActionType(v)
    }

    return fmt.Errorf("can't parse %v int value", value)
}
