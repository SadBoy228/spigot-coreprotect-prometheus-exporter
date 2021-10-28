package modeltypes

import (
    "fmt"
    "database/sql/driver"
)

type BoolTinyInt bool

func (bti BoolTinyInt) Value() (driver.Value, error) {
    if bool(bti) {
        return 1, nil
    }

    return 0, nil
}

func (bti *BoolTinyInt) Scan(value interface{}) error {
    v, ok := value.(int)

    if ok {
        *bti = v == 1
    }

    return fmt.Errorf("can't parse %v value", value)
}
