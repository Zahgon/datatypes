package datatypes

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Time is time data type.
type Time time.Duration

// NewTime is a constructor for Time and returns new Time.
func NewTime(hour, min, sec, nsec int) Time { _ = "STUB: not implemented"; return *new(Time) }

func newTime(hour, min, sec, nsec int) Time { _ = "STUB: not implemented"; return *new(Time) }

// GormDataType returns gorm common data type. This type is used for the field's column type.
func (Time) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType returns gorm DB data type based on the current using database.
	return ""
}

func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

// Scan implements sql.Scanner interface and scans value into Time,
func (t *Time) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (t *Time) setFromString(str string) { _ = "STUB: not implemented"; return }

func (t *Time) setFromTime(src time.Time) { _ = "STUB: not implemented"; return }

// Value implements driver.Valuer interface and returns string format of Time.
func (t Time) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *

	// String implements fmt.Stringer interface.
	new(driver.Value), nil
}

func (t Time) String() string { _ = "STUB: not implemented"; return "" }

// omit nanoseconds unless any value is specified

func (t Time) hours() int { _ = "STUB: not implemented"; return 0 }

func (t Time) minutes() int { _ = "STUB: not implemented"; return 0 }

func (t Time) seconds() int { _ = "STUB: not implemented"; return 0 }

func (t Time) nanoseconds() int { _ = "STUB: not implemented"; return 0 }

// MarshalJSON implements json.Marshaler to convert Time to json serialization.
func (t Time) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler to deserialize json data.
func (t *Time) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// ignore null
	return nil
}
