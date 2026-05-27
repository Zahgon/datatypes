package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"
	"time"
)

// NullString represents a string that may be null.
// NullString implements the [Scanner] interface so
// it can be used as a scan destination:
//
//	var s NullString
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.String
//	} else {
//	   // NULL value
//	}
type NullString = Null[string]

// NullInt64 represents an int64 that may be null.
// NullInt64 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt64 = Null[int64]

// NullInt32 represents an int32 that may be null.
// NullInt32 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt32 = Null[int32]

// NullInt16 represents an int16 that may be null.
// NullInt16 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt16 = Null[int16]

// NullByte represents a byte that may be null.
// NullByte implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullByte = Null[byte]

// NullFloat64 represents a float64 that may be null.
// NullFloat64 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullFloat64 = Null[float64]

// NullBool represents a bool that may be null.
// NullBool implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullBool = Null[bool]

// NullTime represents a [time.Time] that may be null.
// NullTime implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullTime = Null[time.Time]

// Null represents a value that may be null.
// Null implements the [Scanner] interface so
// it can be used as a scan destination:
//
//	var s Null[string]
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.V
//	} else {
//	   // NULL value
//	}
type Null[T any] struct {
	V     T
	Valid bool
}

func (n *Null[T]) Scan(value any) error { _ = "STUB: not implemented"; return nil }

func (n Null[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// NewNull returns a new, non-null Null.
func NewNull[T any](v T) Null[T] { _ = "STUB: not implemented"; return nil }

var errNilPtr = errors.New("destination pointer is nil") // embedded in descriptive error

// convertAssign is the same as convertAssignRows, but without the optional
// rows argument.
func convertAssign(dest, src any) error { _ = "STUB: not implemented"; return nil }

// convertAssignRows copies to dest the value in src, converting it if possible.
// An error is returned if the copy would result in loss of information.
// dest should be a pointer type. If rows is passed in, the rows will
// be used as the parent for any cursor values converted from a
// driver.Rows to a *Rows.
func convertAssignRows(dest, src any, rows *sql.Rows) error {
	_ = "STUB: not implemented"
	// Common cases, without reflect.
	return nil
}

// The following conversions use a string value as an intermediate representation
// to convert between various numeric types.
//
// This also allows scanning into user defined types such as "type Int int64".
// For symmetry, also check for string destination types.

func strconvErr(err error) error { _ = "STUB: not implemented"; return nil }

func asString(src any) string { _ = "STUB: not implemented"; return "" }

func asBytes(buf []byte, rv reflect.Value) (b []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type decimalDecompose interface {
	// Decompose returns the internal decimal state in parts.
	// If the provided buf has sufficient capacity, buf may be returned as the coefficient with
	// the value set and length set as appropriate.
	Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32)
}

type decimalCompose interface {
	// Compose sets the internal decimal value from parts. If the value cannot be
	// represented then an error should be returned.
	Compose(form byte, negative bool, coefficient []byte, exponent int32) error
}

// cloneBytes returns a copy of b[:len(b)].
// The result may have additional unused capacity.
// cloneBytes(nil) returns nil.
func cloneBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }
