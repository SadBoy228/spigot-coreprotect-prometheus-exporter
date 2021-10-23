package modeltypes

import (
    "database/sql/driver"
)

const (
    Logout SessionActionType = iota
    Login
)

type SessionActionType int64

func (s SessionActionType) Value() (driver.Value, error) {
    return int64(s), nil
}

func (s *SessionActionType) Scan(value interface{}) error {
    action, ok := value.(int64)

    if ok {
        *s = SessionActionType(action)
        return nil
    }

    return fmt.Errorf("can't parse %s int value", action)
}
