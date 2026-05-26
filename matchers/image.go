package matchers

var (
	TypeJpeg     = newType("jpg", "image/jpeg")
	TypeJpeg2000 = newType("jp2", "image/jp2")
	TypePng      = newType("png", "image/png")
	TypeGif      = newType("gif", "image/gif")
	TypeWebp     = newType("webp", "image/webp")
	TypeCR2      = newType("cr2", "image/x-canon-cr2")
	TypeTiff     = newType("tif", "image/tiff")
	TypeBmp      = newType("bmp", "image/bmp")
	TypeJxr      = newType("jxr", "image/vnd.ms-photo")
	TypePsd      = newType("psd", "image/vnd.adobe.photoshop")
	TypeIco      = newType("ico", "image/vnd.microsoft.icon")
	TypeHeif     = newType("heif", "image/heif")
	TypeDwg      = newType("dwg", "image/vnd.dwg")
	TypeExr      = newType("exr", "image/x-exr")
	TypeAvif     = newType("avif", "image/avif")
)

var Image = Map{
	TypeJpeg:     Jpeg,
	TypeJpeg2000: Jpeg2000,
	TypePng:      Png,
	TypeGif:      Gif,
	TypeWebp:     Webp,
	TypeCR2:      CR2,
	TypeTiff:     Tiff,
	TypeBmp:      Bmp,
	TypeJxr:      Jxr,
	TypePsd:      Psd,
	TypeIco:      Ico,
	TypeHeif:     Heif,
	TypeDwg:      Dwg,
	TypeExr:      Exr,
	TypeAvif:     Avif,
}

func Jpeg(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Jpeg2000(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Png(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Gif(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Webp(buf []byte) bool { _ = "STUB: not implemented"; return false }

func CR2(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Little Endian
// Big Endian
// CR2 magic word
// CR2 major version

func Tiff(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Little Endian
// Big Endian
// To avoid conflicts differentiate Tiff from CR2

func Bmp(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Jxr(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Psd(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Ico(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Heif(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Dwg(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Exr(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Avif(buf []byte) bool { _ = "STUB: not implemented"; return false }
