package position

// Analyzer checks a single assertion about single char
type Analyzer func(source []byte, start, end int) bool

// True is always `true`
func True(source []byte, start, end int) bool {
	return true
}

// False is always `false`
func False(source []byte, start, end int) bool {
	return false
}

// IsSeparator detects JSON separators: new line chars and spaces
func IsSeparator(source []byte, start, end int) bool {
	char := source[end]
	return char == ' ' || char == '\n' || char == '\t'
}

func IsNumberStart(source []byte, start, end int) bool {
	char := source[end]
	return char == '+' || char == '-' || (char >= '0' && char <= '9')
}

func IsNumberPart(source []byte, start, end int) bool {
	char := source[end]
	return char >= '0' && char <= '9'
}

// IsOpenCurlyBracket looks for `{`
var IsOpenCurlyBracket = isCharAnalyzer('{')

var IsLowerCaseN = isCharAnalyzer('n')

var IsComma = isCharAnalyzer(',')

var IsColon = isCharAnalyzer(':')

func IsLiteralNull(source []byte, start, end int) bool {
	if end-start == 3 {
		if source[start] == 'n' && source[start+1] == 'u' && source[start+2] == 'l' && source[start+3] == 'l' {
			return true
		} else {
			panic("invalid lexeme") // TODO normal error handling
		}
	} else {
		return false
	}
}

// IsCloseCurlyBracket looks for `}`
var IsCloseCurlyBracket = isCharAnalyzer('}')

var IsQuote = isCharAnalyzer('"')

var IsString = func(source []byte, start, end int) bool {
	return start != end && source[end] == '"'
}

func isCharAnalyzer(fixedChar byte) Analyzer {
	return func(source []byte, start, end int) bool {
		return fixedChar == source[end]
	}
}
