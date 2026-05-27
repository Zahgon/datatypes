package datatypes

import (
	"context"
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// JSONMap defined JSON data type, need to implements driver.Valuer, sql.Scanner interface
type JSONMap map[string]interface{}

// Value return json value, implement driver.Valuer interface
func (m JSONMap) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (m *JSONMap) Scan(val interface{}) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON to output non base64 encoded []byte
func (m JSONMap) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON to deserialize []byte
func (m *JSONMap) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// GormDataType gorm common data type
func (m JSONMap) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType gorm db data type
	return ""
}

func (JSONMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (jm JSONMap) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	_ = "STUB: not implemented"
	return *new(clause.Expr)
}
