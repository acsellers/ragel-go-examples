package atof

import (
	"testing"
)

func TestAtof(t *testing.T) {
	for _, test := range atofTests {
		v, e := Atof(test.Data)
		if e != nil && !test.Errors {
			t.Fatalf("Errored on %s", test.Data)
		}
		if v != test.Result {
			t.Fatalf("%v != %v", test.Result, v)
		}
	}
}

type atofTest struct {
	Data   string
	Result float64
	Errors bool
}

var atofTests = []atofTest{
	atofTest{"7", 7.0, false},
	atofTest{"666", 666.0, false},
	atofTest{"-666", -666.0, false},
	atofTest{"+666", 666.0, false},
	atofTest{"1234567890", 1234567890.0, false},
	atofTest{"+1234567890\n", 1234567890.0, false},
	atofTest{"+ 1234567890", 0.0, true},
	atofTest{"0.6", 0.6, false},
	atofTest{"-0.6", -0.6, false},
	atofTest{"-0.1024", -0.1024, false},
	atofTest{"-0.1.024", 0.0, true},
}
