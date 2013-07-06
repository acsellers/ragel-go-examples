
// line 1 "decode.rl"
package hexdecode

import (
  "fmt"
)

func decode(c uint8) uint64 {
  switch c {
  case 'a', 'b', 'c', 'd', 'e', 'f':
    return uint64(c - 'a' + 10)
  case 'A', 'B', 'C', 'D', 'E', 'F':
    return uint64(c - 'A' + 10)
  }
  return 0
}


// line 21 "decode.go"
const hexd_start int = 1
const hexd_first_final int = 3
const hexd_error int = 0

const hexd_en_main int = 1


// line 20 "decode.rl"

func HexDecode(data string) (uint64, error) {
  cs, p, pe := 0, 0, len(data)
  var val uint64
  
// line 35 "decode.go"
	{
	cs = hexd_start
	}

// line 40 "decode.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 2:
		goto st_case_2
	}
	goto st_out
	st_case_1:
		if data[p] == 48 {
			goto tr0
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
// line 27 "decode.rl"

 val = val * 16 + (uint64(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
// line 91 "decode.go"
		switch data[p] {
		case 10:
			goto st4
		case 120:
			goto st2
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		goto st0
tr2:
// line 27 "decode.rl"

 val = val * 16 + (uint64(data[p]) - '0') 
	goto st5
tr3:
// line 28 "decode.rl"

 val = val * 16 + decode(data[p]) 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
// line 132 "decode.go"
		if data[p] == 10 {
			goto st4
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	st_out:
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 33 "decode.rl"

  if cs < hexd_first_final {
    return 0, fmt.Errorf("hexdecode parse error: %s", data)
  }

  return val, nil
}
