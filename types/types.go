package types

import "sync"

// Types Support concurrent map writes
var Types sync.Map

// Add registers a new type in the package
func Add(t Type) Type { _ = "STUB: not implemented"; return *new(Type) }

// Get retrieves a Type by extension
func Get(ext string) Type { _ = "STUB: not implemented"; return *new(Type) }
