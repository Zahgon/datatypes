package datatypes

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// This datatype stores the uuid in the database as a string. To store the uuid
// in the database as a binary (byte) array, please refer to datatypes.BinUUID.
type UUID uuid.UUID

// NewUUIDv1 generates a UUID version 1, panics on generation failure.
func NewUUIDv1() UUID { _ = "STUB: not implemented"; return *new(UUID) }

// NewUUIDv4 generates a UUID version 4, panics on generation failure.
func NewUUIDv4() UUID { _ = "STUB: not implemented"; return *new(UUID) }

// NewUUIDv7 generates a UUID version 7, panics on generation failure.
func NewUUIDv7() UUID { _ = "STUB: not implemented"; return *new(UUID) }

// GormDataType gorm common data type.
func (UUID) GormDataType() string {
	_ = "STUB: not implemented"

	// GormDBDataType gorm db data type.
	return ""
}

func (UUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

// Scan is the scanner function for this datatype.
func (u *UUID) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Value is the valuer function for this datatype.
func (u UUID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// String returns the string form of the UUID.
func (u UUID) String() string { _ = "STUB: not implemented"; return "" }

// Equals returns true if string form of UUID matches other, false otherwise.
func (u UUID) Equals(other UUID) bool { _ = "STUB: not implemented"; return false }

// Length returns the number of characters in string form of UUID.
func (u UUID) Length() int { _ = "STUB: not implemented"; return 0 }

// IsNil returns true if the UUID is a nil UUID (all zeroes), false otherwise.
func (u UUID) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if UUID is nil UUID or of zero length, false otherwise.
func (u UUID) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsNilPtr returns true if caller UUID ptr is nil, false otherwise.
func (u *UUID) IsNilPtr() bool {
	_ = "STUB: not implemented"

	// IsEmptyPtr returns true if caller UUID ptr is nil or it's value is empty.
	return false
}

func (u *UUID) IsEmptyPtr() bool { _ = "STUB: not implemented"; return false }
