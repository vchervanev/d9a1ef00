package bjson

import "testing"

func TestEmpty(t *testing.T) {
	Parse([]byte("{}"))
}
func TestEmptyAndSeparators(t *testing.T) {
	Parse([]byte("  {  }  "))
}

func TestString(t *testing.T) {
	Parse([]byte("{\"dd\":\"value\"}"))
}

func TestStringAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : \"value\" } "))
}

func TestNumberAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : 0 } "))
}
func TestNumber(t *testing.T) {
	Parse([]byte("{\"dd\":0}"))
}

func TestNegativeNumberAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : -1 } "))
}
func TestNegativeNumber(t *testing.T) {
	Parse([]byte("{\"dd\":-1}"))
}

func TestPositiveNumberAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : +1 } "))
}
func TestPositiveNumber(t *testing.T) {
	Parse([]byte("{\"dd\":+1}"))
}

func TestLongNumberAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : +12345 } "))
}
func TestLongNumber(t *testing.T) {
	Parse([]byte("{\"dd\":+12345}"))
}

func TestNullAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : null } "))
}
func TestNull(t *testing.T) {
	Parse([]byte("{\"dd\":null}"))
}

func TestPairAndSeparators(t *testing.T) {
	Parse([]byte(" { \"dd\" : null,\n\"ff\": \"QQQ\" } "))
}

func TestSampleData(t *testing.T) {
	const json = "{\n		\"id\": 123,\n		\"email\": \"robosen@icloud.com\",\n		\"first_name\": \"Данила\",\n		\"last_name\": \"Стамленский\",\n		\"gender\": \"m\",\n		\"birth_date\": -345081600\n}"
	Parse([]byte(json))
}
