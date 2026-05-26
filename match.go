package filetype

import (
	"io"

	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"
)

// Matchers is an alias to matchers.Matchers
var Matchers = matchers.Matchers

// MatcherKeys is an alias to matchers.MatcherKeys
var MatcherKeys = &matchers.MatcherKeys

// NewMatcher is an alias to matchers.NewMatcher
var NewMatcher = matchers.NewMatcher

// Match infers the file type of a given buffer inspecting its magic numbers signature
func Match(buf []byte) (types.Type, error) { _ = "STUB: not implemented"; return *new(types.Type), nil }

// Get is an alias to Match()
func Get(buf []byte) (types.Type, error) {
	_ = "STUB: not implemented"

	// MatchFile infers a file type for a file
	return *new(types.Type), nil
}

func MatchFile(filepath string) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

// MatchReader is convenient wrapper to Match() any Reader
func MatchReader(reader io.Reader) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

// 8K makes msooxml tests happy and allows for expanded custom file checks

// AddMatcher registers a new matcher type
func AddMatcher(fileType types.Type, matcher matchers.Matcher) matchers.TypeMatcher {
	_ = "STUB: not implemented"
	return *new(matchers.TypeMatcher)
}

// Matches checks if the given buffer matches with some supported file type
func Matches(buf []byte) bool { _ = "STUB: not implemented"; return false }

// MatchMap performs a file matching against a map of match functions
func MatchMap(buf []byte, matchers matchers.Map) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

// MatchesMap is an alias to Matches() but using matching against a map of match functions
func MatchesMap(buf []byte, matchers matchers.Map) bool { _ = "STUB: not implemented"; return false }
