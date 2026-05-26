package types

// Type represents a file MIME type and its extension
type Type struct {
	MIME      MIME
	Extension string
}

// NewType creates a new Type
func NewType(ext, mime string) Type { _ = "STUB: not implemented"; return *new(Type) }
