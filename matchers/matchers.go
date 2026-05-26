package matchers

import (
	"github.com/h2non/filetype/types"
)

// Internal shortcut to NewType
var newType = types.NewType

// Matcher function interface as type alias
type Matcher func([]byte) bool

// Type interface to store pairs of type with its matcher function
type Map map[types.Type]Matcher

// Type specific matcher function interface
type TypeMatcher func([]byte) types.Type

// Store registered file type matchers
var Matchers = make(map[types.Type]TypeMatcher)
var MatcherKeys []types.Type

// Create and register a new type matcher function
func NewMatcher(kind types.Type, fn Matcher) TypeMatcher {
	_ = "STUB: not implemented"
	return *new(TypeMatcher)
}

// prepend here so any user defined matchers get added first

func register(matchers ...Map) { _ = "STUB: not implemented"; return }

func init() {
	// Arguments order is intentional
	// Archive files will be checked last due to prepend above in func NewMatcher
	register(Archive, Document, Font, Audio, Video, Image, Application)
}
