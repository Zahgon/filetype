package matchers

var (
	TypeDoc  = newType("doc", "application/msword")
	TypeDocx = newType("docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	TypeXls  = newType("xls", "application/vnd.ms-excel")
	TypeXlsx = newType("xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	TypePpt  = newType("ppt", "application/vnd.ms-powerpoint")
	TypePptx = newType("pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation")
	TypeOdp  = newType("odp", "application/vnd.oasis.opendocument.presentation")
	TypeOds  = newType("ods", "application/vnd.oasis.opendocument.spreadsheet")
	TypeOdt  = newType("odt", "application/vnd.oasis.opendocument.text")
)

var Document = Map{
	TypeDoc:  Doc,
	TypeDocx: Docx,
	TypeXls:  Xls,
	TypeXlsx: Xlsx,
	TypePpt:  Ppt,
	TypePptx: Pptx,
	TypeOdp:  Odp,
	TypeOds:  Ods,
	TypeOdt:  Odt,
}

type docType int

const (
	TYPE_DOC docType = iota
	TYPE_DOCX
	TYPE_XLS
	TYPE_XLSX
	TYPE_PPT
	TYPE_PPTX
	TYPE_OOXML
	TYPE_ODP
	TYPE_ODS
	TYPE_ODT
)

// reference: https://bz.apache.org/ooo/show_bug.cgi?id=111457
func Doc(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Docx(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Xls(buf []byte) bool { _ = "STUB: not implemented"; return false }

// BIFF5 && BIFF12(12)

// BIFF12(11)

func Xlsx(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Ppt(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Pptx(buf []byte) bool { _ = "STUB: not implemented"; return false }

func msooxml(buf []byte) (typ docType, found bool) {
	_ = "STUB: not implemented"
	return *new(docType), false
}

// start by checking for ZIP local file header signature

// make sure the first file is correct

// skip to the second local file header
// since some documents include a 520-byte extra field following the file
// header, we need to scan for the next header

// now skip to the *third* local file header; again, we need to scan due to a
// 520-byte extra field following the file header

// and check the subdirectory name to determine which type of OOXML
// file we have.  Correct the mimetype with the registered ones:
// http://technet.microsoft.com/en-us/library/cc179224.aspx

// OpenOffice/Libreoffice orders ZIP entry differently, so check the 4th file

func compareBytes(slice, subSlice []byte, startOffset int) bool {
	_ = "STUB: not implemented"
	return false
}

func checkMSOoml(buf []byte, offset int) (typ docType, ok bool) {
	_ = "STUB: not implemented"
	return *new(docType), false
}

func search(buf []byte, start, rangeNum int) int { _ = "STUB: not implemented"; return 0 }

func Odp(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Ods(buf []byte) bool { _ = "STUB: not implemented"; return false }

func Odt(buf []byte) bool { _ = "STUB: not implemented"; return false }

// https://en.wikipedia.org/wiki/OpenDocument_technical_specification
// https://en.wikipedia.org/wiki/ZIP_(file_format)
func checkOdf(buf []byte, mimetype string) bool { _ = "STUB: not implemented"; return false }

// Perform all byte checks first for better performance
// Check ZIP start

// Now check the first file data
// Compression method: not compressed

// Filename length must be 8 for "mimetype"

// Check the file contents sizes

// No extra field (for data offset below)

// Finally check the file name and contents
