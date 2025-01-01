package core

import (
  "bytes"
  "testing"
)

func TestParagraphsShouldBeInterruptedByBlankLine(t *testing.T) {
  var tests = []struct{
    input string
    want  string
  } {
    {"aaa\n\nbbb", "Paragraph: aaa\nParagraph: bbb\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestParagraphsShouldBeInterruptedByBlankLine(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}
