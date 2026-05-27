package datatypes

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// This datatype is similar to datatypes.UUID, major difference being that
// this datatype stores the uuid in the database as a binary (byte) array
// instead of a string. Developers may use either as per their preference.
type BinUUID uuid.UUID

// NewBinUUIDv1 generates a uuid version 1, panics on generation failure.
func NewBinUUIDv1() BinUUID { _ = "STUB: not implemented"; return *new(BinUUID) }

// NewBinUUIDv4 generates a uuid version 4, panics on generation failure.
func NewBinUUIDv4() BinUUID { _ = "STUB: not implemented"; return *new(BinUUID) }

// NewNilBinUUID generates a nil uuid.
func NewNilBinUUID() BinUUID {
	_ = "STUB: not implemented"
	return *

	// BinUUIDFromString returns the BinUUID representation of the specified uuidStr.
	new(BinUUID)
}

func BinUUIDFromString(uuidStr string) BinUUID { _ = "STUB: not implemented"; return *new(BinUUID) }

// GormDataType gorm common data type.
func (BinUUID) GormDataType() string { _ = "STUB: not implemented"; return "" }

// GormDBDataType gorm db data type.
func (BinUUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	_ = "STUB: not implemented"
	return ""
}

// Scan is the scanner function for this datatype.
func (u *BinUUID) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Value is the valuer function for this datatype.
func (u BinUUID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// String returns the string form of the UUID.
func (u BinUUID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the string form of the UUID.
func (u BinUUID) String() string { _ = "STUB: not implemented"; return "" }

// Equals returns true if bytes form of BinUUID matches other, false otherwise.
func (u BinUUID) Equals(other BinUUID) bool { _ = "STUB: not implemented"; return false }

// Length returns the number of characters in string form of UUID.
func (u BinUUID) LengthBytes() int { _ = "STUB: not implemented"; return 0 }

// Length returns the number of characters in string form of UUID.
func (u BinUUID) Length() int { _ = "STUB: not implemented"; return 0 }

// IsNil returns true if the BinUUID is nil uuid (all zeroes), false otherwise.
func (u BinUUID) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if BinUUID is nil uuid or of zero length, false otherwise.
func (u BinUUID) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsNilPtr returns true if caller BinUUID ptr is nil, false otherwise.
func (u *BinUUID) IsNilPtr() bool {
	_ = "STUB: not implemented"

	// IsEmptyPtr returns true if caller BinUUID ptr is nil or it's value is empty.
	return false
}

func (u *BinUUID) IsEmptyPtr() bool { _ = "STUB: not implemented"; return false }
