package datatypes

import (
	"context"
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// JSONType give a generic data type for json encoded data.
type JSONType[T any] struct {
	data T
}

func NewJSONType[T any](data T) JSONType[T] { _ = "STUB: not implemented"; return nil }

// Data return data with generic Type T
func (j JSONType[T]) Data() T {
	_ = "STUB: not implemented"

	// Value return json value, implement driver.Valuer interface
	return *new(T)
}

func (j JSONType[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *JSONType[T]) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON to output non base64 encoded []byte
func (j JSONType[T]) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalJSON to deserialize []byte
		nil
}

func (j *JSONType[T]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// GormDataType gorm common data type
func (JSONType[T]) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType gorm db data type
	return ""
}

func (JSONType[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (js JSONType[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	_ = "STUB: not implemented"
	return *new(clause.Expr)
}

// JSONSlice give a generic data type for json encoded slice data.
type JSONSlice[T any] []T

func NewJSONSlice[T any](s []T) JSONSlice[T] { _ = "STUB: not implemented"; return nil }

// Value return json value, implement driver.Valuer interface
func (j JSONSlice[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *JSONSlice[T]) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// GormDataType gorm common data type
func (JSONSlice[T]) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType gorm db data type
	return ""
}

func (JSONSlice[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (j JSONSlice[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	_ = "STUB: not implemented"
	return *new(clause.Expr)
}
