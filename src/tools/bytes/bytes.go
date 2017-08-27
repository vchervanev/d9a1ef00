package bytes

func Eql(bytes1, bytes2 []byte) bool {
	if len(bytes1) != len(bytes2) {
		return false
	}
	for i, b := range bytes1 {
		if b != bytes2[i] {
			return false
		}
	}
	return true
}

func StartsWith(bytes, prefix []byte) bool {
	for i, b := range prefix {
		if b != bytes[i] {
			return false
		}
	}
	return true
}

func GetId(bytes []byte, separator byte) int {
	result := 0
	power := 1
	i := len(bytes) - 1
	for ; bytes[i] != separator; i -= 1 {
		result += int(bytes[i]-'0') * power
		power *= 10
	}
	return result
}

func IndexOfSubarray(source []byte, value []byte) int {
	max_start := len(source) - len(value)
	for i := 0; i < max_start; i++ {
		equal := true
		for j := 0; j < len(value); j++ {
			if source[i+j] != value[j] {
				equal = false
				break
			}
		}

		if equal {
			return i
		}
	}
	return -1
}

// TODO use indexed prefixes
func IndexOf(bytes [][]byte, value []byte) int {
	for i, candidate := range bytes {
		if len(candidate) != len(value) {
			continue
		}
		equal := true
		for j, v := range candidate {
			if v != value[j] {
				equal = false
				break
			}
		}
		if equal {
			return i
		}
	}
	return -1
}
