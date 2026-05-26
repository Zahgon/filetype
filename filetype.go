package filetype

import (
	"errors"

	"github.com/h2non/filetype/types"
)

// Types stores a map of supported types
var Types = types.Types

// NewType creates and registers a new type
var NewType = types.NewType

// Unknown represents an unknown file type
var Unknown = types.Unknown

// ErrEmptyBuffer represents an empty buffer error
var ErrEmptyBuffer = errors.New("Empty buffer")

// ErrUnknownBuffer represents a unknown buffer error
var ErrUnknownBuffer = errors.New("Unknown buffer type")

// AddType registers a new file type
func AddType(ext, mime string) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

// Is checks if a given buffer matches with the given file type extension
func Is(buf []byte, ext string) bool { _ = "STUB: not implemented"; return false }

// IsExtension semantic alias to Is()
func IsExtension(buf []byte, ext string) bool { _ = "STUB: not implemented"; return false }

// IsType checks if a given buffer matches with the given file type
func IsType(buf []byte, kind types.Type) bool { _ = "STUB: not implemented"; return false }

// IsMIME checks if a given buffer matches with the given MIME type
func IsMIME(buf []byte, mime string) bool { _ = "STUB: not implemented"; return false }

// IsSupported checks if a given file extension is supported
func IsSupported(ext string) bool { _ = "STUB: not implemented"; return false }

// IsMIMESupported checks if a given MIME type is supported
func IsMIMESupported(mime string) bool { _ = "STUB: not implemented"; return false }

// GetType retrieves a Type by file extension
func GetType(ext string) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }
