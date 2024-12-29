package core

import (
	"testing"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestHeadersMustHaveOneToSixHashesAtBeggining(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Token
	}{
		{"# Heading", blocks.NewHeaderToken("Heading", 1)},
		{"## Heading", blocks.NewHeaderToken("Heading", 2)},
		{"### Heading", blocks.NewHeaderToken("Heading", 3)},
		{"#### Heading", blocks.NewHeaderToken("Heading", 4)},
		{"##### Heading", blocks.NewHeaderToken("Heading", 5)},
		{"###### Heading", blocks.NewHeaderToken("Heading", 6)},
    {"####### Heading", blocks.NewParagraphBlock("####### Heading")},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestTokenizesHeaders(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

func TestHeadersMustHaveSpaceOrTabAfterHashes(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Token
	}{
		{"#Heading", blocks.NewParagraphBlock("#Heading")},
		{"# Heading", blocks.NewHeaderToken("Heading", 1)},
		{"##    Heading", blocks.NewHeaderToken("Heading", 2)},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestHeadersMustHaveSpaceOrTabAfterHashes(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

func TestHeadersFirstHashMustNotBeEscaped(t *testing.T) {
  var tests = []struct {
    input string
    want blocks.Token
  }{
    {"\\# Heading", blocks.NewParagraphBlock("\\# Heading")},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestHeadersFirstHashMustNotBeEscaped(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestSpacesAndTabsAtBeggingAndEndingOfHeadingsAreIgnored(t *testing.T) {
  var tests = []struct {
    input string
    want blocks.Token
  }{
    {"#      Heading        ", blocks.NewHeaderToken("Heading", 1)},
    {"##                Heading", blocks.NewHeaderToken("Heading", 2)},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestSpacesAndTabsAtBegginingAndEndingOfHeadingsAreIgnored(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestUpToThreeSpacesOfIdentationAreAllowedOnHeadings(t *testing.T) {
  var tests = []struct {
    input string
    want blocks.Token
  }{
    {" ### Heading", blocks.NewHeaderToken("Heading", 3)},
    {"  #### Heading", blocks.NewHeaderToken("Heading", 4)},
    {"   ###### Heading", blocks.NewHeaderToken("Heading", 6)},
    {"    # Heading", blocks.NewParagraphBlock("    # Heading")},
  }
  
  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestUpToThreeSpacesOfIdentationAreAllowedOnHeadings(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}
