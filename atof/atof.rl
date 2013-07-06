package atof

import (
  "fmt"
)

%%{
  machine atof;
  write data;
}%%
func Atof(data string) (float64, error) {
  cs, p, pe := 0, 0, len(data)
  var neg bool
  var val float64
  lowness := float64(10.0)
  %%{
    action see_neg { neg = true }
    action add_whole { val = val * 10 + (float64(fc) - '0') }
    action add_decimal { val = val + ((float64(fc) - '0') / lowness); lowness = lowness * 10.0; }

    main :=
      ( '-'@see_neg | '+')? ( digit @add_whole ) + ('.' ( digit @add_decimal )* )? '\n'?;
    write init;
    write exec;
  }%%
  if neg {
    val = -val
  }
  if cs < atof_first_final {
    return 0.0, fmt.Errorf("atof parse error: %s", data)
  }

  return val, nil
}
