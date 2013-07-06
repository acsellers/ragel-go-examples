
// line 1 "atoi.rl"
package atoi

import (
  "fmt"
)


// line 11 "atoi.go"
const atoi_start int = 1
const atoi_first_final int = 3
const atoi_error int = 0

const atoi_en_main int = 1


// line 10 "atoi.rl"


func Atoi(data string) (int64, error) {
  cs, p, pe := 0, 0, len(data)
  var neg bool
  var val int64
  
// line 27 "atoi.go"
	{
	cs = atoi_start
	}

// line 32 "atoi.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 43:
			goto st2
		case 45:
			goto tr2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
// line 17 "atoi.rl"

 neg = true 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
// line 75 "atoi.go"
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
tr3:
// line 18 "atoi.rl"

 val = val * 10 + (int64(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
// line 90 "atoi.go"
		if data[p] == 10 {
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 24 "atoi.rl"


  if neg {
    val = -val
  }

  if cs < atoi_first_final {
    return 0, fmt.Errorf("atoi parse error: %s", data)
  }

  return val, nil
}
