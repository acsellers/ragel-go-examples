package comma

import (
  "fmt"
)

%%{
  machine comma;
  write data;
}%%

func Comma(data string) ([]string, error) {
  cs, p, pe := 0, 0, len(data)
  val := make([]string, 0)
  var current string

  %%{
    noncomma = [^,];
    normal = noncomma & [^"];

    action separator { val = append(val, current); current = "" }
    action character { current = current + string(fc) }
    action escapedCharacter { current = current + "\\" + fc }

    bare = ( normal @character )* ( ',' @separator );
    surrounded = '"' ( normal @character )* '"' + ( ',' @separator );
    nonterminator = bare | surrounded;
    main := ( nonterminator )* ( normal @character )*;

    write init;
    write exec;
  }%%
  val = append(val, current)

  if cs < comma_first_final {
    return []string{}, fmt.Errorf("comma parse error: %s", data)
  }

  return val, nil
}
