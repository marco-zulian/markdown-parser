package core

import (
  "testing"
)

func TestShouldDetectCodeBlocks(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"mock_files/code.md", "Code: a simple\n  indented code block\n"}, 
  }

  for _, test := range tests {
    if result, _ := Tokenize(test.input); *result.GetContent() != test.want {
      t.Errorf("TestShouldDetectCodeBlocks(%s) = \n%q, want \n%q", test.input, *result.GetContent(), test.want)
    }
  }
}
