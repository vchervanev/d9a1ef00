package bjson

import "testing"

type hash map[string]string

func assertParsing(input string, expected hash, t *testing.T) {
	names, values := Parse([]byte(input))
	for index, bName := range names {
		name := string(bName)
		value := string(values[index])
		expectedValue := expected[name]

		if value != expectedValue {
			t.Error(
				"For", name,
				"expected", expectedValue,
				"got", value,
			)

		}
	}
	if len(names) != len(values) {
		t.Error(
			"Odd names/values combination", len(names), len(values),
		)
	}
	if len(names) != len(expected) {
		t.Error(
			"Odd attributes/ expected values combination", len(names), len(expected),
		)
	}
}

func TestEmpty(t *testing.T) {
	assertParsing("{}", nil, t)
}
func TestEmptyAndSeparators(t *testing.T) {
	assertParsing(" { } ", nil, t)
}

func TestString(t *testing.T) {
	assertParsing("{\"dd\":\"value\"}", hash{"dd": "value"}, t)
}

func TestStringAndSeparators(t *testing.T) {
	assertParsing(" { \"dd\" : \"value\" } ", hash{"dd": "value"}, t)
}

func TestNumberAndSeparators(t *testing.T) {
	assertParsing(" { \"dd\" : 0 } ", hash{"dd": "0"}, t)
}
func TestNumber(t *testing.T) {
	assertParsing("{\"dd\":0}", hash{"dd": "0"}, t)
}

func TestNegativeNumberAndSeparators(t *testing.T) {
	assertParsing(" { \"dd\" : -1 } ", hash{"dd": "-1"}, t)
}
func TestNegativeNumber(t *testing.T) {
	assertParsing("{\"dd\":-1}", hash{"dd": "-1"}, t)
}

func TestPositiveNumberAndSeparators(t *testing.T) {
	assertParsing(" { \"dd\" : +1 } ", hash{"dd": "+1"}, t)
}
func TestPositiveNumber(t *testing.T) {
	assertParsing("{\"dd\":+1}", hash{"dd": "+1"}, t)
}

func TestLongNumberAndSeparators(t *testing.T) {
	assertParsing(" { \"dd\" : +12345 } ", hash{"dd": "+12345"}, t)
}
func TestLongNumber(t *testing.T) {
	assertParsing("{\"dd\":+12345}", hash{"dd": "+12345"}, t)
}

func assertDDNull(names, values [][]byte, t *testing.T) {
	if len(names) != 1 || len(values) != 1 {
		t.Error("Invalid response length")
	}
	if string(names[0]) != "dd" {
		t.Error("Invalid attr name, expected dd, got", names[0])
	}
	if len(values[0]) != 0 {
		t.Error("Invalid attr value, expected nil, got", string(values[0]))
	}
}

func TestNullAndSeparators(t *testing.T) {
	names, values := Parse([]byte(" { \"dd\" : null } "))
	assertDDNull(names, values, t)
}
func TestNull(t *testing.T) {
	names, values := Parse([]byte("{\"dd\":null}"))
	assertDDNull(names, values, t)
}

func TestPairAndSeparators(t *testing.T) {
	names, values := Parse([]byte(" { \"dd\" : null,\n\"ff\": \"QQQ\" } "))
	assertDDNull(names[0:1], values[0:1], t)
	if string(names[1]) != "ff" || string(values[1]) != "QQQ" {
		t.Error("another error")
	}
}

func TestSampleData(t *testing.T) {
	const json = "{\n		\"id\": 123,\n		\"email\": \"robosen@icloud.com\",\n		\"first_name\": \"Данила\",\n		\"last_name\": \"Стамленский\",\n		\"gender\": \"m\",\n		\"birth_date\": -345081600\n}"
	var expected = hash{
		"id":         "123",
		"email":      "robosen@icloud.com",
		"first_name": "Данила",
		"last_name":  "Стамленский",
		"gender":     "m",
		"birth_date": "-345081600",
	}
	assertParsing(json, expected, t)
}
