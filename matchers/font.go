package matchers

var (
	TypeWoff  = newType("woff", "application/font-woff")
	TypeWoff2 = newType("woff2", "application/font-woff")
	TypeTtf   = newType("ttf", "application/font-sfnt")
	TypeOtf   = newType("otf", "application/font-sfnt")
)

var Font = Map{
	TypeWoff:  Woff,
	TypeWoff2: Woff2,
	TypeTtf:   Ttf,
	TypeOtf:   Otf,
}

func Woff(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Woff2(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Ttf(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Otf(buf []byte) bool { _ = "STUB: not implemented"; return false }
