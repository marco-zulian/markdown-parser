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
    {"aaa\nbbb\n\nccc\nddd", "Paragraph: aaa\nbbb\nParagraph: ccc\nddd\n"},
    {"aaa\n\n\nbbb", "Paragraph: aaa\nParagraph: bbb\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestParagraphsShouldBeInterruptedByBlankLine(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestLeadingSpacesAndTabsInParagraphLinesAreTrimmed(t *testing.T) {
  var tests = []struct{
    input string
    want  string
  } {
    {"  aaa\n bbb", "Paragraph: aaa\nbbb\n"},
    {"aaa\n             bbb\n                                       ccc", "Paragraph: aaa\nbbb\nccc\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestLeadingSpacesAndTabsInParagraphLinesAreTrimmed(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}
