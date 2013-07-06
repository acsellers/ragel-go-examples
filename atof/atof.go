// line 1 "atof.rl"
package atof

import (
	"fmt"
)

// line 12 "atof.go"
const atof_start int = 1
const atof_first_final int = 3
const atof_error int = 0

const atof_en_main int = 1

// line 11 "atof.rl"

func Atof(data string) (float64, error) {
	cs, p, pe := 0, 0, len(data)
	var neg bool
	var val float64
	lowness := float64(10.0)

	// line 28 "atof.go"
	{
		cs = atof_start
	}

	// line 33 "atof.go"
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
		case 5:
			goto st_case_5
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
		// line 18 "atof.rl"

		neg = true
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		// line 78 "atof.go"
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
	tr3:
		// line 19 "atof.rl"

		val = val*10 + (float64(data[p]) - '0')
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		// line 93 "atof.go"
		switch data[p] {
		case 10:
			goto st4
		case 46:
			goto st5
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
	tr6:
		// line 20 "atof.rl"

		val = val + ((float64(data[p]) - '0') / lowness)
		lowness = lowness * 10.0
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		// line 120 "atof.go"
		if data[p] == 10 {
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	// line 26 "atof.rl"

	if neg {
		val = -val
	}
	if cs < atof_first_final {
		return 0.0, fmt.Errorf("atof parse error: %s", data)
	}

	return val, nil
}
