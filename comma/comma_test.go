package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	for _, test := range commaTests {
		v, e := Comma(test.Data)
		if e != nil && !test.Errors {
			t.Fatalf("Errored on %s", test.Data)
		}
		if !strEqual(v, test.Result) {
			t.Fatalf("%v != %v", v, test.Result)
		}
	}
}

func strEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

type commaTest struct {
	Data   string
	Result []string
	Errors bool
}

var commaTests = []commaTest{
	commaTest{"a,b,c", []string{"a", "b", "c"}, false},
	commaTest{"", []string{""}, false},
	commaTest{"a,,c", []string{"a", "", "c"}, false},
	commaTest{"1,2,3", []string{"1", "2", "3"}, false},
	commaTest{"11,22,33", []string{"11", "22", "33"}, false},
	commaTest{"\"1\",2", []string{"1", "2"}, false},
}
