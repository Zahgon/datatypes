package datatypes

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// JSON defined JSON data type, need to implements driver.Valuer, sql.Scanner interface
type JSON json.RawMessage

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON to output non base64 encoded []byte
func (j JSON) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON to deserialize []byte
func (j *JSON) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (j JSON) String() string {
	_ = "STUB: not implemented"

	// GormDataType gorm common data type
	return ""
}

func (JSON) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType gorm db data type
	return ""
}

func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (js JSON) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	_ = "STUB: not implemented"
	return *new(clause.Expr)
}

// JSONQueryExpression json query expression, implements clause.Expression interface to use as querier
type JSONQueryExpression struct {
	column      string
	keys        []string
	hasKeys     bool
	equals      bool
	likes       bool
	equalsValue interface{}
	extract     bool
	path        string
}

// JSONQuery query column as json
func JSONQuery(column string) *JSONQueryExpression { _ = "STUB: not implemented"; return nil }

// Extract extract json with path
func (jsonQuery *JSONQueryExpression) Extract(path string) *JSONQueryExpression {
	_ = "STUB: not implemented"
	return nil
}

// HasKey returns clause.Expression
func (jsonQuery *JSONQueryExpression) HasKey(keys ...string) *JSONQueryExpression {
	_ = "STUB: not implemented"
	return nil
}

// Keys returns clause.Expression
func (jsonQuery *JSONQueryExpression) Equals(value interface{}, keys ...string) *JSONQueryExpression {
	_ = "STUB: not implemented"
	return nil
}

// Likes return clause.Expression
func (jsonQuery *JSONQueryExpression) Likes(value interface{}, keys ...string) *JSONQueryExpression {
	_ = "STUB: not implemented"
	return nil
}

// Build implements clause.Expression
func (jsonQuery *JSONQueryExpression) Build(builder clause.Builder) {
	_ = "STUB: not implemented"
	return
}

// JSONOverlapsExpression JSON_OVERLAPS expression, implements clause.Expression interface to use as querier
type JSONOverlapsExpression struct {
	column clause.Expression
	val    string
}

// JSONOverlaps query column as json
func JSONOverlaps(column clause.Expression, value string) *JSONOverlapsExpression {
	_ = "STUB: not implemented"
	return nil
}

// Build implements clause.Expression
// only mysql support JSON_OVERLAPS
func (json *JSONOverlapsExpression) Build(builder clause.Builder) {
	_ = "STUB: not implemented"
	return
}

type columnExpression string

func Column(col string) columnExpression { _ = "STUB: not implemented"; return *new(columnExpression) }

func (col columnExpression) Build(builder clause.Builder) { _ = "STUB: not implemented"; return }

const prefix = "$."

func jsonQueryJoin(keys []string) string { _ = "STUB: not implemented"; return "" }

// JSONSetExpression json set expression, implements clause.Expression interface to use as updater
type JSONSetExpression struct {
	column     string
	path2value map[string]interface{}
	mutex      sync.RWMutex
}

// JSONSet update fields of json column
func JSONSet(column string) *JSONSetExpression { _ = "STUB: not implemented"; return nil }

// Set return clause.Expression.
//
//	{
//		"age": 20,
//		"name": "json-1",
//		"orgs": {"orga": "orgv"},
//		"tags": ["tag1", "tag2"]
//	}
//
//	// In MySQL/SQLite, path is `age`, `name`, `orgs.orga`, `tags[0]`, `tags[1]`.
//	DB.UpdateColumn("attr", JSONSet("attr").Set("orgs.orga", 42))
//
//	// In PostgreSQL, path is `{age}`, `{name}`, `{orgs,orga}`, `{tags, 0}`, `{tags, 1}`.
//	DB.UpdateColumn("attr", JSONSet("attr").Set("{orgs, orga}", "bar"))
func (jsonSet *JSONSetExpression) Set(path string, value interface{}) *JSONSetExpression {
	_ = "STUB: not implemented"
	return nil
}

// Build implements clause.Expression
// support mysql, sqlite and postgres
func (jsonSet *JSONSetExpression) Build(builder clause.Builder) { _ = "STUB: not implemented"; return }

func JSONArrayQuery(column string) *JSONArrayExpression { _ = "STUB: not implemented"; return nil }

type JSONArrayExpression struct {
	contains    bool
	in          bool
	column      string
	keys        []string
	equalsValue interface{}
}

// Contains checks if column[keys] contains the value given. The keys parameter is only supported for MySQL and SQLite.
func (json *JSONArrayExpression) Contains(value interface{}, keys ...string) *JSONArrayExpression {
	_ = "STUB: not implemented"
	return nil
}

// In checks if columns[keys] is in the array value given. This method is only supported for MySQL and SQLite.
func (json *JSONArrayExpression) In(value interface{}, keys ...string) *JSONArrayExpression {
	_ = "STUB: not implemented"
	return nil
}

// Build implements clause.Expression
func (json *JSONArrayExpression) Build(builder clause.Builder) { _ = "STUB: not implemented"; return }
