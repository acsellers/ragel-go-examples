package hexdecode

import (
	"testing"
)

func TestHex(t *testing.T) {
	for _, test := range hexTests {
		v, e := HexDecode(test.Data)
		if e != nil && !test.Errors {
			t.Fatalf("Errored on %s", test.Data)
		}
		if v != test.Result {
			t.Fatalf("%v != %v", test.Result, v)
		}
	}

}

type hexTest struct {
	Data   string
	Result uint64
	Errors bool
}

var hexTests = []hexTest{
	hexTest{"7", 7, false},
	hexTest{"0x7", 7, false},
	hexTest{"0xa", 10, false},
	hexTest{"0xA", 10, false},
	hexTest{"1A", 26, false},
	hexTest{"0x2A", 42, false},
}
