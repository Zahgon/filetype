package filetype

import (
	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"
)

// Image tries to match a file as image type
func Image(buf []byte) (types.Type, error) { _ = "STUB: not implemented"; return *new(types.Type), nil }

// IsImage checks if the given buffer is an image type
func IsImage(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Audio tries to match a file as audio type
func Audio(buf []byte) (types.Type, error) { _ = "STUB: not implemented"; return *new(types.Type), nil }

// IsAudio checks if the given buffer is an audio type
func IsAudio(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Video tries to match a file as video type
func Video(buf []byte) (types.Type, error) { _ = "STUB: not implemented"; return *new(types.Type), nil }

// IsVideo checks if the given buffer is a video type
func IsVideo(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Font tries to match a file as text font type
func Font(buf []byte) (types.Type, error) { _ = "STUB: not implemented"; return *new(types.Type), nil }

// IsFont checks if the given buffer is a font type
func IsFont(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Archive tries to match a file as generic archive type
func Archive(buf []byte) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

// IsArchive checks if the given buffer is an archive type
func IsArchive(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Document tries to match a file as document type
func Document(buf []byte) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

// IsDocument checks if the given buffer is an document type
func IsDocument(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Application tries to match a file as an application type
func Application(buf []byte) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

// IsApplication checks if the given buffer is an application type
func IsApplication(buf []byte) bool { _ = "STUB: not implemented"; return false }

func doMatchMap(buf []byte, machers matchers.Map) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}
