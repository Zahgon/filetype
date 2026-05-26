package isobmff

// IsISOBMFF checks whether the given buffer represents ISO Base Media File Format data
func IsISOBMFF(buf []byte) bool { _ = "STUB: not implemented"; return false }

// GetFtyp returns the major brand, minor version and compatible brands of the ISO-BMFF data
func GetFtyp(buf []byte) (string, string, []string) { _ = "STUB: not implemented"; return "", "", nil }
