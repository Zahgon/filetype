package matchers

var (
	TypeMp4  = newType("mp4", "video/mp4")
	TypeM4v  = newType("m4v", "video/x-m4v")
	TypeMkv  = newType("mkv", "video/x-matroska")
	TypeWebm = newType("webm", "video/webm")
	TypeMov  = newType("mov", "video/quicktime")
	TypeAvi  = newType("avi", "video/x-msvideo")
	TypeWmv  = newType("wmv", "video/x-ms-wmv")
	TypeMpeg = newType("mpg", "video/mpeg")
	TypeFlv  = newType("flv", "video/x-flv")
	Type3gp  = newType("3gp", "video/3gpp")
)

var Video = Map{
	TypeMp4:  Mp4,
	TypeM4v:  M4v,
	TypeMkv:  Mkv,
	TypeWebm: Webm,
	TypeMov:  Mov,
	TypeAvi:  Avi,
	TypeWmv:  Wmv,
	TypeMpeg: Mpeg,
	TypeFlv:  Flv,
	Type3gp:  Match3gp,
}

func M4v(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Mkv(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Webm(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Mov(buf []byte) bool { _ = "STUB: not implemented"; return false }

// 'f' 't'
// 'y' 'p'
// 'q' 't'

func Avi(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Wmv(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Mpeg(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Flv(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Mp4(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Match3gp(buf []byte) bool { _ = "STUB: not implemented"; return false }

func containsMatroskaSignature(buf, subType []byte) bool { _ = "STUB: not implemented"; return false }
