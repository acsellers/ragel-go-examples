
// line 1 "comma.rl"
package comma

import (
  "fmt"
)


// line 11 "comma.go"
const comma_start int = 3
const comma_first_final int = 3
const comma_error int = 0

const comma_en_main int = 3


// line 10 "comma.rl"


func Comma(data string) ([]string, error) {
  cs, p, pe := 0, 0, len(data)
  val := make([]string, 0)
  var current string

  
// line 28 "comma.go"
	{
	cs = comma_start
	}

// line 33 "comma.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	}
	goto st_out
tr3:
// line 21 "comma.rl"

 val = append(val, current); current = "" 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
// line 61 "comma.go"
		switch data[p] {
		case 34:
			goto st1
		case 44:
			goto tr3
		}
		goto tr4
tr4:
// line 22 "comma.rl"

 current = current + string(data[p]) 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
// line 79 "comma.go"
		switch data[p] {
		case 34:
			goto st0
		case 44:
			goto tr3
		}
		goto tr4
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
// line 22 "comma.rl"

 current = current + string(data[p]) 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
// line 101 "comma.go"
		switch data[p] {
		case 34:
			goto st2
		case 44:
			goto st0
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 34:
			goto st2
		case 44:
			goto tr3
		}
		goto st0
	st_out:
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 32 "comma.rl"

  val = append(val, current)

  if cs < comma_first_final {
    return []string{}, fmt.Errorf("comma parse error: %s", data)
  }

  return val, nil
}
