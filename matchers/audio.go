package matchers

var (
	TypeMidi = newType("mid", "audio/midi")
	TypeMp3  = newType("mp3", "audio/mpeg")
	TypeM4a  = newType("m4a", "audio/mp4")
	TypeOgg  = newType("ogg", "audio/ogg")
	TypeFlac = newType("flac", "audio/x-flac")
	TypeWav  = newType("wav", "audio/x-wav")
	TypeAmr  = newType("amr", "audio/amr")
	TypeAac  = newType("aac", "audio/aac")
	TypeAiff = newType("aiff", "audio/x-aiff")
)

var Audio = Map{
	TypeMidi: Midi,
	TypeMp3:  Mp3,
	TypeM4a:  M4a,
	TypeOgg:  Ogg,
	TypeFlac: Flac,
	TypeWav:  Wav,
	TypeAmr:  Amr,
	TypeAac:  Aac,
	TypeAiff: Aiff,
}

func Midi(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Mp3(buf []byte) bool { _ = "STUB: not implemented"; return false }

func M4a(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Ogg(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Flac(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Wav(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Amr(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Aac(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Aiff(buf []byte) bool { _ = "STUB: not implemented"; return false }
