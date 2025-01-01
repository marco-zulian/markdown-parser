package core

import (
  "bytes"
  "testing"
)

func TestCodeFencesBlockBeginWithEitherBackticksOrTildes(t *testing.T) {
  var tests = []struct{
    input string
    want  string
  }{
    {"```\n<\n >\n```", "Fenced code: <\n >\n"},
    {"~~~\n<\n >\n~~~", "Fenced code: <\n >\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestCodeFencesBlockBeginWithEitherBackticksOrTildes(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  } 

}
