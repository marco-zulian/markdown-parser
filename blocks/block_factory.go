package blocks

import (
  "regexp"
  "strings"
)

var blockTypeRegexs = map[BlockType]*regexp.Regexp{
  Header        : regexp.MustCompile(`^( {0,3})(#{1,6})([ \t]+|$)`),
  ThematicBreak : regexp.MustCompile(`^( {0,3})((\*[ \*]*\*[ \*]*\*[ \*]*)|(-[ -]*-[ -]*-[ -]*)|(_[ _]*_[ _]*_[ _]*))$`),
  Code          : regexp.MustCompile(`^ {4,}|\t`),
  BlankLine     : regexp.MustCompile(`^ *$`),
  FencedCode    : regexp.MustCompile("^`{3,}|~{3,}"),
}

func GenerateBlock(line string) Block {
  if blockTypeRegexs[BlankLine].Match([]byte(line)) { return nil }

  var blockTypeConstructors = map[BlockType]func(string) Block {
    Header        : func(line string) Block { return NewHeaderBlock(line) },
    ThematicBreak : func(line string) Block { return NewThematicBreakBlock(line) },
    Code          : func(line string) Block { return NewCodeBlock(line) },
    FencedCode    : func(line string) Block { return NewFencedCodeBlock(line) },
  }

  var processingOrder = []BlockType{Code, Header, ThematicBreak, FencedCode}
  
  for _, blockType := range processingOrder {
    re, _ := blockTypeRegexs[blockType]
    if re.Match([]byte(line)) {
      constructor, _ := blockTypeConstructors[blockType]
      return constructor(line)
    }
  }
  
  return NewParagraphBlock(line)
}

func NewHeaderBlock(line string) *HeaderBlock {
	hashRe := regexp.MustCompile(`#{1,6}`)
  endingRe := regexp.MustCompile(` [# ]+$`)
  headingLevel := len(hashRe.Find([]byte(line)))

  trimmedString := endingRe.ReplaceAllString(strings.TrimLeft(line, "# "), "")
	
  return &HeaderBlock{
		content : trimmedString,
		Level   : headingLevel,
	}
}

func NewCodeBlock(line string) *CodeBlock {
  if strings.HasPrefix(line, "\t") { return &CodeBlock{ content: line[1:], isOpen:  true } 
  } else if strings.HasPrefix(line, " \t") { return &CodeBlock{ content: line[2:], isOpen: true }
  } else if strings.HasPrefix(line, "  \t") { return &CodeBlock{ content: line[3:], isOpen: true }
  } else { return &CodeBlock{ content: line[4:], isOpen:  true } }
}

func NewParagraphBlock(content string) *ParagraphBlock {
	return &ParagraphBlock{
		content: strings.TrimLeft(content, " \t"),
    isOpen: true,
	}
}

func NewThematicBreakBlock(line string) *ThematicBreakBlock {
  return &ThematicBreakBlock{}
}

func NewFencedCodeBlock(line string) *FencedCodeBlock {
  return &FencedCodeBlock {
    content: "",
    delimiter: string(line[0]),
    info: "",
    isOpen: true,
    tabs: 0,
  }
}
