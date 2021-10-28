package modeltypes

import (
    "fmt"

    "database/sql/driver"
)

const (
    SessionLogout SessionActionType = iota
    SessionLogin
)

type SessionActionType int

func (s SessionActionType) Value() (driver.Value, error) {
    return int(s), nil
}

func (s *SessionActionType) Scan(value interface{}) error {
    action, ok := value.(int)

    if ok {
        *s = SessionActionType(action)
        return nil
    }

    return fmt.Errorf("can't parse %v int value", value)
}
