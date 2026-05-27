package datatypes

import (
	"database/sql/driver"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type URL url.URL

func (u URL) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (u *URL) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (URL) GormDataType() string { _ = "STUB: not implemented"; return "" }

func (URL) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (u *URL) String() string { _ = "STUB: not implemented"; return "" }

func (u URL) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (u *URL) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// ignore null
	return nil
}
