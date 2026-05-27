package datatypes

import (
	"database/sql/driver"
	"time"
)

type Date time.Time

func (date *Date) Scan(value interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func (date Date) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string { _ = "STUB: not implemented"; return "" }

func (date Date) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (date *Date) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

func (date Date) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (date *Date) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
