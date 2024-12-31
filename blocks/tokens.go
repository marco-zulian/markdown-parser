package blocks

type BlockType string
const (
	Document      BlockType = "document"
  Paragraph     BlockType = "paragraph"
	Header        BlockType = "header"
	ThematicBreak BlockType = "break"
	Code          BlockType = "code"
  Html          BlockType = "html"
  Link          BlockType = "link"
  BlockQuote    BlockType = "blockquote"
  List          BlockType = "list"
  ListItem      BlockType = "list_item"
  Image         BlockType = "image"
  LineBreak     BlockType = "line_break"
)

type Block interface {
	GetBlockType() BlockType
	GetContent() *string
  CanConsume(line string) bool
  Consume(line string)
  CanExtend() bool
}


