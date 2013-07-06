package atoi

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	for _, test := range atoiTests {
		v, e := Atoi(test.Data)
		if e != nil && !test.Errors {
			t.Fatalf("Errored on %s", test.Data)
		}
		if v != test.Result {
			t.Fatalf("%v != %v", test.Result, v)
		}
	}
}

type atoiTest struct {
	Data   string
	Result int64
	Errors bool
}

var atoiTests = []atoiTest{
	atoiTest{"7", 7, false},
	atoiTest{"666", 666, false},
	atoiTest{"-666", -666, false},
	atoiTest{"+666", 666, false},
	atoiTest{"1234567890", 1234567890, false},
	atoiTest{"+1234567890\n", 1234567890, false},
	atoiTest{"+ 1234567890", 0, true},
}
