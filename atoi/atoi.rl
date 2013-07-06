package atoi

import (
  "fmt"
)

%%{
  machine atoi;
  write data;
}%%

func Atoi(data string) (int64, error) {
  cs, p, pe := 0, 0, len(data)
  var neg bool
  var val int64
  %%{
    action see_neg { neg = true }
    action add_digit { val = val * 10 + (int64(fc) - '0') }

    main :=
      ( '-'@see_neg | '+')? ( digit @add_digit ) + '\n'?;
    write init;
    write exec;
  }%%

  if neg {
    val = -val
  }

  if cs < atoi_first_final {
    return 0, fmt.Errorf("atoi parse error: %s", data)
  }

  return val, nil
}
