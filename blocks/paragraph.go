package blocks

import (
  "fmt"
)

type ParagraphBlock struct {
	content string
  isOpen  bool
}

func (paragraph *ParagraphBlock) GetBlockType() BlockType {
	return Paragraph
}

func (paragraph *ParagraphBlock) GetContent() *string {
	return &paragraph.content
}

func (paragraph *ParagraphBlock) String() string {
  return fmt.Sprintf("Paragraph: %s", paragraph.content)
}

func (paragraph *ParagraphBlock) CanConsume(line string) bool {
  return false
}

func (paragraph *ParagraphBlock) Consume(line string) {
  paragraph.content += "\n" + line 
}

func (paragraph *ParagraphBlock) CanExtend() bool {
  return true 
}
