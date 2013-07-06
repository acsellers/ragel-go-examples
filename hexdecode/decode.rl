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

%%{
  machine hexd;
  write data;
}%%
func HexDecode(data string) (uint64, error) {
  cs, p, pe := 0, 0, len(data)
  var val uint64
  %%{
    hex_alpha = [a-fA-F];

    action add_digit { val = val * 16 + (uint64(fc) - '0') }
    action add_char { val = val * 16 + decode(fc) }

    main := '0x'? ( digit @add_digit | hex_alpha @add_char ) + '\n'?;
    write init;
    write exec;
  }%%
  if cs < hexd_first_final {
    return 0, fmt.Errorf("hexdecode parse error: %s", data)
  }

  return val, nil
}
