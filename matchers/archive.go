package matchers

const (
	ZstdMagicSkippableStart = 0x184D2A50
	ZstdMagicSkippableMask  = 0xFFFFFFF0
)

var (
	TypeEpub    = newType("epub", "application/epub+zip")
	TypeZip     = newType("zip", "application/zip")
	TypeTar     = newType("tar", "application/x-tar")
	TypeRar     = newType("rar", "application/vnd.rar")
	TypeGz      = newType("gz", "application/gzip")
	TypeBz2     = newType("bz2", "application/x-bzip2")
	Type7z      = newType("7z", "application/x-7z-compressed")
	TypeXz      = newType("xz", "application/x-xz")
	TypeZstd    = newType("zst", "application/zstd")
	TypePdf     = newType("pdf", "application/pdf")
	TypeExe     = newType("exe", "application/vnd.microsoft.portable-executable")
	TypeSwf     = newType("swf", "application/x-shockwave-flash")
	TypeRtf     = newType("rtf", "application/rtf")
	TypeEot     = newType("eot", "application/octet-stream")
	TypePs      = newType("ps", "application/postscript")
	TypeSqlite  = newType("sqlite", "application/vnd.sqlite3")
	TypeNes     = newType("nes", "application/x-nintendo-nes-rom")
	TypeCrx     = newType("crx", "application/x-google-chrome-extension")
	TypeCab     = newType("cab", "application/vnd.ms-cab-compressed")
	TypeDeb     = newType("deb", "application/vnd.debian.binary-package")
	TypeAr      = newType("ar", "application/x-unix-archive")
	TypeZ       = newType("Z", "application/x-compress")
	TypeLz      = newType("lz", "application/x-lzip")
	TypeRpm     = newType("rpm", "application/x-rpm")
	TypeElf     = newType("elf", "application/x-executable")
	TypeDcm     = newType("dcm", "application/dicom")
	TypeIso     = newType("iso", "application/x-iso9660-image")
	TypeMachO   = newType("macho", "application/x-mach-binary") // Mach-O binaries have no common extension.
	TypeParquet = newType("parquet", "application/vnd.apache.parquet")
)

var Archive = Map{
	TypeEpub:    bytePrefixMatcher(epubMagic),
	TypeZip:     Zip,
	TypeTar:     Tar,
	TypeRar:     Rar,
	TypeGz:      bytePrefixMatcher(gzMagic),
	TypeBz2:     bytePrefixMatcher(bz2Magic),
	Type7z:      bytePrefixMatcher(sevenzMagic),
	TypeXz:      bytePrefixMatcher(xzMagic),
	TypeZstd:    Zst,
	TypePdf:     bytePrefixMatcher(pdfMagic),
	TypeExe:     bytePrefixMatcher(exeMagic),
	TypeSwf:     Swf,
	TypeRtf:     bytePrefixMatcher(rtfMagic),
	TypeEot:     Eot,
	TypePs:      bytePrefixMatcher(psMagic),
	TypeSqlite:  bytePrefixMatcher(sqliteMagic),
	TypeNes:     bytePrefixMatcher(nesMagic),
	TypeCrx:     bytePrefixMatcher(crxMagic),
	TypeCab:     Cab,
	TypeDeb:     bytePrefixMatcher(debMagic),
	TypeAr:      bytePrefixMatcher(arMagic),
	TypeZ:       Z,
	TypeLz:      bytePrefixMatcher(lzMagic),
	TypeRpm:     Rpm,
	TypeElf:     Elf,
	TypeDcm:     Dcm,
	TypeIso:     Iso,
	TypeMachO:   MachO,
	TypeParquet: bytePrefixMatcher(parquetMagic),
}

var (
	epubMagic = []byte{
		0x50, 0x4B, 0x03, 0x04, 0x6D, 0x69, 0x6D, 0x65,
		0x74, 0x79, 0x70, 0x65, 0x61, 0x70, 0x70, 0x6C,
		0x69, 0x63, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x2F,
		0x65, 0x70, 0x75, 0x62, 0x2B, 0x7A, 0x69, 0x70,
	}
	gzMagic     = []byte{0x1F, 0x8B, 0x08}
	bz2Magic    = []byte{0x42, 0x5A, 0x68}
	sevenzMagic = []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C}
	pdfMagic    = []byte{0x25, 0x50, 0x44, 0x46}
	exeMagic    = []byte{0x4D, 0x5A}
	rtfMagic    = []byte{0x7B, 0x5C, 0x72, 0x74, 0x66}
	nesMagic    = []byte{0x4E, 0x45, 0x53, 0x1A}
	crxMagic    = []byte{0x43, 0x72, 0x32, 0x34}
	psMagic     = []byte{0x25, 0x21}
	xzMagic     = []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00}
	sqliteMagic = []byte{0x53, 0x51, 0x4C, 0x69}
	debMagic    = []byte{
		0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E, 0x0A,
		0x64, 0x65, 0x62, 0x69, 0x61, 0x6E, 0x2D, 0x62,
		0x69, 0x6E, 0x61, 0x72, 0x79,
	}
	arMagic      = []byte{0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E}
	zstdMagic    = []byte{0x28, 0xB5, 0x2F, 0xFD}
	lzMagic      = []byte{0x4C, 0x5A, 0x49, 0x50}
	parquetMagic = []byte{0x50, 0x41, 0x52, 0x31}
)

func bytePrefixMatcher(magicPattern []byte) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

func Zip(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Tar(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Rar(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Swf(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Cab(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Eot(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Z(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Rpm(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Elf(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Dcm(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Iso(buf []byte) bool { _ = "STUB: not implemented"; return false }

func MachO(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Big endian versions below here...

// Zstandard compressed data is made of one or more frames.
// There are two frame formats defined by Zstandard: Zstandard frames and Skippable frames.
// See more details from https://tools.ietf.org/id/draft-kucherawy-dispatch-zstd-00.html#rfc.section.2
func Zst(buf []byte) bool { _ = "STUB: not implemented"; return false }

// skippable frames
