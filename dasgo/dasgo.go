// Package dasgo provides Type Definitions for the Discord API.
package dasgo

// Snowflake represents a Discord API Snowflake.
type Snowflake uint64

// Flag represents an (unused) alias for a Discord API Flag ranging from 0 - 255.
type Flag uint8

// BitFlag represents an alias for a Discord API Bitwise Flag denoted by 1 << x.
type BitFlag uint64

// File represents a file attachment.
type File struct {
	Name        string
	ContentType string
	Data        []byte
}

// Nonce represents a Discord nonce (integer or string).
type Nonce interface{}

// Value represents a value (string, integer, or double).
type Value interface{}

// Pointer returns a pointer to the given value.
func Pointer[T any](v T) *T {
	return &v
}
