// Package dasgo provides Type Definitions for the Discord API.
package dasgo

// Snowflake represents a Discord API Snowflake.
type Snowflake uint64

// Flag represents an alias for a Discord API Flag ranging from 0 - 255.
type Flag uint8

// Flags represents an alias for a []Flag for JSON marshal purposes.
type Flags []Flag

// BitFlag represents an alias for a Discord API Bitwise Flag denoted by 1 << x.
type BitFlag uint64

// File represents a file attachment.
type File struct {
	Name        string
	ContentType string
	Data        []byte
}

// Timestamp represents an ISO8601 timestamp.
type Timestamp string

// Nonce represents a Discord nonce (integer or string).
type Nonce interface{}

// Value represents a value (string, integer, double, or bool).
type Value interface{}

// Pointer returns a pointer to the given value.
func Pointer[T any](v T) *T {
	return &v
}

// Pointer2 returns a double pointer to the given value.
//
// set `null` to true in order to point the double pointer to a `nil` pointer.
func Pointer2[T any](v T, null ...bool) **T {
	if len(null) != 0 && null[0] {
		return new(*T)
	}

	pointer := Pointer(v)

	return &pointer
}

// IsValue returns whether the given pointer contains a value.
func IsValue[T any](p *T) bool {
	return p != nil
}

// IsValue2 returns whether the given double pointer contains a pointer.
func IsValue2[T any](dp **T) bool {
	return dp != nil
}

// PointerIndicator represents a Dasgo double pointer value indicator.
type PointerIndicator uint8

const (
	// IsValueNothing indicates that the field was not provided.
	//
	// The double pointer is nil.
	IsValueNothing PointerIndicator = 0

	// IsValueNull indicates the field was provided with a null value.
	//
	// The double pointer points to a nil pointer.
	IsValueNull PointerIndicator = 1

	// IsValueValid indicates that the field is a valid value.
	//
	// The double pointer points to a pointer that points to a value.
	IsValueValid PointerIndicator = 2
)

// PointerCheck returns whether the given double pointer contains a value.
//
// returns IsValueNothing, IsValueNull, or IsValueValid.
//
//	IsValueNothing indicates that the field was not provided.
//	IsValueNull indicates the field was provided with a null value.
//	IsValueValid indicates that the field is a valid value.
func PointerCheck[T any](dp **T) PointerIndicator {
	if dp != nil {
		if *dp != nil {
			return IsValueValid
		}

		return IsValueNull
	}

	return IsValueNothing
}
