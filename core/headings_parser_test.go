package core

import (
	"testing"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestHeadersMustHaveOneToSixHashesAtBeggining(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Block
	}{
		{"# Heading", blocks.NewHeaderBlock("Heading", 1)},
		{"## Heading", blocks.NewHeaderBlock("Heading", 2)},
		{"### Heading", blocks.NewHeaderBlock("Heading", 3)},
		{"#### Heading", blocks.NewHeaderBlock("Heading", 4)},
		{"##### Heading", blocks.NewHeaderBlock("Heading", 5)},
		{"###### Heading", blocks.NewHeaderBlock("Heading", 6)},
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
		want  blocks.Block
	}{
		{"#Heading", blocks.NewParagraphBlock("#Heading")},
		{"# Heading", blocks.NewHeaderBlock("Heading", 1)},
		{"##    Heading", blocks.NewHeaderBlock("Heading", 2)},
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
		want  blocks.Block
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
		want  blocks.Block
	}{
		{"#      Heading        ", blocks.NewHeaderBlock("Heading", 1)},
		{"##                Heading", blocks.NewHeaderBlock("Heading", 2)},
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
		want  blocks.Block
	}{
		{" ### Heading", blocks.NewHeaderBlock("Heading", 3)},
		{"  #### Heading", blocks.NewHeaderBlock("Heading", 4)},
		{"   ###### Heading", blocks.NewHeaderBlock("Heading", 6)},
		{"    # Heading", blocks.NewCodeBlock("# Heading")},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestUpToThreeSpacesOfIdentationAreAllowedOnHeadings(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

func TestHeadingsMightBeEmpty(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Block
	}{
		{"## ", blocks.NewHeaderBlock("", 2)},
		{"#", blocks.NewHeaderBlock("", 1)},
		{"### ###", blocks.NewHeaderBlock("", 3)},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestHeadingMightBeEmpty(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

func TestHeadingsClosingSequencesAreIgnored(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Block
	}{
		{"## foo ##", blocks.NewHeaderBlock("foo", 2)},
		{"  ###   bar    ###", blocks.NewHeaderBlock("bar", 3)},
		{"# foo ##################################", blocks.NewHeaderBlock("foo", 1)},
		{"##### foo ##", blocks.NewHeaderBlock("foo", 5)},
		{"### foo ###     ", blocks.NewHeaderBlock("foo", 3)},
		{"### foo ### b", blocks.NewHeaderBlock("foo ### b", 3)},
		{"# foo#", blocks.NewHeaderBlock("foo#", 1)},
		{"### foo \\###", blocks.NewHeaderBlock("foo \\###", 3)},
		{"## foo #\\##", blocks.NewHeaderBlock("foo #\\##", 2)},
		{"# foo \\#", blocks.NewHeaderBlock("foo \\#", 1)},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestHeadingsClosingSequencesAreIgnored(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}
