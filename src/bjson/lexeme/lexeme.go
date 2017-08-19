package lexeme

import "../position"
import "../value"

type Lexeme struct {
	Id      byte
	Name    string
	IsStart position.Analyzer // detects starting symbol
	IsPart  position.Analyzer // detects internal symbols // false means the lexeme was finished at prev char
	IsDone  position.Analyzer // detects final symbol (only if its IsPart == true)

	GetValue value.Reader // return value, e.g. remove quotetion marks
}

func GetNext(lex Lexeme, json []byte, pos int) Lexeme {
	for _, next := range transitions[lex.Id] {
		if next.IsStart(json, pos, pos) {
			return next
		}
	}
	panic("unable to parse json")
}

// pseudo lexeme before JSON body
var Outside = Lexeme{
	Id:      0,
	Name:    "outside",
	IsStart: position.True,
	IsPart:  position.False,
}

var start = Lexeme{
	Id:      1,
	Name:    "start",
	IsStart: position.IsOpenCurlyBracket,
	IsPart:  position.False,
	IsDone:  position.True,
}

var attributeName = Lexeme{
	Id:      2,
	Name:    "attributeName",
	IsStart: position.IsQuote,
	IsPart:  position.True,
	IsDone:  position.IsString,

	GetValue: value.String,
}

var attributeSeparator = Lexeme{
	Id:      3,
	Name:    "attributeSeparator",
	IsStart: position.IsColon,
	IsPart:  position.False,
	IsDone:  position.True,
}

var attributeValueString = Lexeme{
	Id:      4,
	Name:    "attributeValueString",
	IsStart: position.IsQuote,
	IsPart:  position.True,
	IsDone:  position.IsString,

	GetValue: value.String,
}

var attributeValueNull = Lexeme{
	Id:      5,
	Name:    "attributeValueNull",
	IsStart: position.IsLowerCaseN,
	IsPart:  position.True,
	IsDone:  position.IsLiteralNull,

	GetValue: value.Constant,
}

var attributeValueNumber = Lexeme{
	Id:      6,
	Name:    "attributeValueNumber",
	IsStart: position.IsNumberStart,
	IsPart:  position.IsNumberPart,
	IsDone:  position.False,

	GetValue: value.Constant,
}

var attributePairSeparator = Lexeme{
	Id:      7,
	Name:    "attributePairSeparator",
	IsStart: position.IsComma,
	IsPart:  position.False,
	IsDone:  position.True,
}

var End = Lexeme{
	Id:      8,
	Name:    "end",
	IsStart: position.IsCloseCurlyBracket,
	IsPart:  position.False,
}

var transitions = [][]Lexeme{
	// outside
	{start},
	// start
	{attributeName, End},
	// attributeName
	{attributeSeparator},
	// attributeSeparator
	{attributeValueString, attributeValueNull, attributeValueNumber},
	// attributeValueString
	{attributePairSeparator, End},
	// attributeValueNull
	{attributePairSeparator, End},
	// attributeValueNumber
	{attributePairSeparator, End},
	// attributePairSeparator
	{attributeName},
	// end
	{},
}
