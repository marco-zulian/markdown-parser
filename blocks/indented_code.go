package blocks

import (
  "fmt"
)

type CodeBlock struct {
	content string
  isOpen  bool
}

func (codeBlock CodeBlock) GetBlockType() BlockType {
	return Code
}

func (codeBlock CodeBlock) GetContent() *string {
	return &codeBlock.content
}

func (codeBlock CodeBlock) String() string {
  return fmt.Sprintf("Code: %s", codeBlock.content)
}

func (codeBlock *CodeBlock) CanConsume(line string) bool {
  return false
}

func (codeBlock *CodeBlock) Consume(line string) {
  codeBlock.content += "\n" + line[4:] 
}

func (codeBlock *CodeBlock) CanExtend() bool {
  return true 
}

