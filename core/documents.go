package core

import (
  "fmt"
  "regexp"

  "github.com/marco-zulian/markdown-parser/blocks"
)

type Document struct {
  blocks  []blocks.Block
  isOpen  bool
}

func NewDocument() Document {
  return Document{
    blocks: []blocks.Block{},
    isOpen: true,
  }
}

func (document Document) GetBlockType() blocks.BlockType {
  return blocks.Document
}

func (document Document) GetContent() *string {
  var content string

  for _, block := range document.blocks {
    content += fmt.Sprintf("%s\n", block)
  }

  return &content
}

func (document *Document) IngestLine(line string) {
  if len(document.blocks) == 0 { 
    document.blocks = append(document.blocks, blocks.GenerateBlock(line)) 
    return
  }

  openBlock := document.blocks[len(document.blocks) - 1]
  
  if openBlock.CanConsume(line) {
    openBlock.Consume(line)
    return
  }

  if !openBlock.CanExtend() || isNewBlockStart(line) {
    document.blocks = append(document.blocks, blocks.GenerateBlock(line)) 
  } else {
    openBlock.Consume(line)
  }

}

func isNewBlockStart(line string) bool {
  var blockTypeRegexs = map[blocks.BlockType]*regexp.Regexp{}

  var processingOrder = []blocks.BlockType{}
  
  for _, blockType := range processingOrder {
    re, _ := blockTypeRegexs[blockType]
    if re.Match([]byte(line)) { 
      return true
    }
  }

  return false
}
