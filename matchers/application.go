package matchers

var (
	TypeWasm = newType("wasm", "application/wasm")
	TypeDex  = newType("dex", "application/vnd.android.dex")
	TypeDey  = newType("dey", "application/vnd.android.dey")
)

var Application = Map{
	TypeWasm: Wasm,
	TypeDex:  Dex,
	TypeDey:  Dey,
}

// Wasm detects a Web Assembly 1.0 filetype.
func Wasm(buf []byte) bool {
	_ = "STUB: not implemented"
	// WASM has starts with `\0asm`, followed by the version.
	// http://webassembly.github.io/spec/core/binary/modules.html#binary-magic
	return false
}

// Dex detects dalvik executable(DEX)
func Dex(buf []byte) bool {
	_ = "STUB: not implemented"
	// https://source.android.com/devices/tech/dalvik/dex-format#dex-file-magic
	return false
}

// magic

// file sise

// Dey Optimized Dalvik Executable(ODEX)
func Dey(buf []byte) bool { _ = "STUB: not implemented"; return false }

// dey magic

// dex
