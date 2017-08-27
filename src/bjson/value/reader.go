package value

// Reader parses lexeme's value
type Reader func(source []byte) []byte

// Null returns nil (empty slice?)
func Null(source []byte) []byte { return nil }

// String slices `source` to cut off starting and ending quotes
func String(source []byte) []byte { return dup(source[1 : len(source)-1]) }

// Constant returns `source` value itself
func Constant(source []byte) []byte { return dup(source) }

func dup(src []byte) (result []byte) {
	result = make([]byte, len(src))
	copy(result, src)
	return
}
