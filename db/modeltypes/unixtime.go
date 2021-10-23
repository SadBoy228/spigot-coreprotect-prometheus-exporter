package modeltypes

import (
    "fmt"
    "time"

    "database/sql/driver"
)

type Unixtime struct{
    time.Time
}

func (u Unixtime) Value() (driver.Value, error) {
    return u.Time.Unix(), nil
}

func (u *Unixtime) Scan(value interface{}) error {
    unixtime, ok := value.(int64)

    if ok {
        *u = Unixtime{Time: time.Unix(unixtime, 0)}
        return nil
    }

    return fmt.Errorf("Can't parse %v time value: ", u)
}
