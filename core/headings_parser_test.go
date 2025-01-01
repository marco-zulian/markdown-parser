package core

import (
	"testing"
  "bytes"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestHeadersMustHaveOneToSixHashesAtBeggining(t *testing.T) {
	var tests = []struct {
		input string
		want  string 
	}{
      {"# Heading", "Header 1: Heading"},
      {"## Heading", "Header 2: Heading"},
      {"### Heading", "Header 3: Heading"},
      {"#### Heading", "Header 4: Heading"},
      {"##### Heading", "Header 5: Heading"},
      {"###### Heading", "Header 6: Heading"},
	    {"####### Heading", "Paragraph: ####### Heading"},
  }

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestTokenizesHeaders(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestHeadersMustHaveSpaceOrTabAfterHashes(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {"#Heading", "Paragraph: #Heading"},
      {"# Heading", "Header 1: Heading"},
      {"##    Heading", "Header 2: Heading"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestHeadersMustHaveSpaceOrTabAfterHashes(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestHeadersMustHaveSpaceOrTabAfterHashesMultiline(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"#5 bolt\n\n#hashtag", "Paragraph: #5 bolt\nParagraph: #hashtag\n"},
    {"foo\n    # bar", "Paragraph: foo\n# bar\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
			t.Errorf("TestHeadersMustHaveSpaceOrTabAfterHashesMultiline(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestHeadersFirstHashMustNotBeEscaped(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {"\\# Heading", "Paragraph: \\# Heading"},
      {"\\## foo", "Paragraph: \\## foo"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestHeadersFirstHashMustNotBeEscaped(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestSpacesAndTabsAtBeggingAndEndingOfHeadingsAreIgnored(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {"#      Heading        ", "Header 1: Heading"},
      {"##                Heading", "Header 2: Heading"},
	    {"#                  foo                     ", "Header 1: foo"},  
  }

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestSpacesAndTabsAtBegginingAndEndingOfHeadingsAreIgnored(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestUpToThreeSpacesOfIdentationAreAllowedOnHeadings(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {" ### foo", "Header 3: foo"},
      {"  ## foo", "Header 2: foo"},
      {"   # foo", "Header 1: foo"},
      {"    # foo", "Code: # foo"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestUpToThreeSpacesOfIdentationAreAllowedOnHeadings(%s) = %s, want %s", test.input, result.String(), test.want)
		}
	}
}

func TestHeadingsMightBeEmpty(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
    {"## ", "Header 2: "},
    {"#", "Header 1: "},
    {"### ###", "Header 3: "},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestHeadingMightBeEmpty(%s) = %s, want %s", test.input, result.String(), test.want)
		}
	}
}

func TestHeadingsClosingSequencesAreIgnored(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {"## foo ##", "Header 2: foo"},
      {"  ###   bar    ###", "Header 3: bar"},
      {"# foo ##################################", "Header 1: foo"},
      {"##### foo ##", "Header 5: foo"},
      {"### foo ###     ", "Header 3: foo"},
      {"### foo ### b", "Header 3: foo ### b"},
      {"# foo#", "Header 1: foo#"},
      {"### foo \\###", "Header 3: foo \\###"},
      {"## foo #\\##", "Header 2: foo #\\##"},
      {"# foo \\#", "Header 1: foo \\#"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestHeadingsClosingSequencesAreIgnored(%s) = %s, want %s", test.input, result.String(), test.want)
		}
	}
}

func TestHeadingsDontNeedToBeSeparatedFromSurroundingContentByBlankLine(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"****\n## foo\n****", "Thematic break\nHeader 2: foo\nThematic break\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
			t.Errorf("TestHeadingsClosingSequencesMultiline(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestHeadingsCanInterruptParagraphs(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"Foo bar\n# baz\nBar foo", "Paragraph: Foo bar\nHeader 1: baz\nParagraph: Bar foo\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestHeadingsCanInterruptParagraphs(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}
