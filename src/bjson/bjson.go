package bjson

import (
	"./lexeme"
	"./position"
)

// var log = fmt.Printf
func log(format string, a ...interface{}) {}

func Parse(json []byte) (names [][]byte, values [][]byte) {
	start := 0
	pos := 0
	lex := lexeme.Outside
	active := false
	for pos < len(json) {
		log("Pos = %v, active = %v, id = %v\n", pos, active, lex.Name)
		if active {
			isPart := lex.IsPart(json, start, pos)
			isDone := isPart && lex.IsDone(json, start, pos)
			log("isPart = %v, isDone = %v\n", isPart, isDone)
			if !isPart || isDone {
				log("Deactivated\n")
				active = false
				if lex.GetValue != nil {
					delta := 0
					if isPart {
						delta = 1
					}
					value := lex.GetValue(json[start : pos+delta])
					log("%v\n\n", string(value))

					if lexeme.IsAttribute(lex) {
						names = append(names, value)
					} else {
						values = append(values, value)
					}
				}
			}
			if !isPart && !position.IsSeparator(json, pos, pos) {
				// duplication!
				start = pos
				lex = lexeme.GetNext(lex, json, pos)
				active = lex.IsDone == nil || !lex.IsDone(json, pos, pos)
				log("next %v : active = %v \n", lex.Name, active)
			}
		} else {
			// inactive - skip separators or find next lex
			if position.IsSeparator(json, pos, pos) {
				log("Inactive - skipped separator\n")
				// nothing
			} else {
				start = pos
				lex = lexeme.GetNext(lex, json, pos)
				active = lex.IsDone == nil || !lex.IsDone(json, pos, pos)
				log("next %v : active = %v \n", lex.Name, active)
			}
		}

		pos = pos + 1
	}
	if lex.Id != lexeme.End.Id {
		panic("invalid end lexeme")
	}

	return
}

// ToBytes transforms raw JSON (as []byte) into fixed-size byte array (or [][]byte?)
// func ToBytes(json []byte, attributes []string, sizes []int) (result []byte) {
// }

// Build builds raw JSON (as []byte) from bytes array and type definition
// func Build(bjson []byte, attributes []string, sizes []int) (json []byte) {

// }
