package blocks

type BlockType string
const (
	Document      BlockType = "document"
  Paragraph     BlockType = "paragraph"
	Header        BlockType = "header"
	ThematicBreak BlockType = "break"
	Code          BlockType = "code"
  FencedCode    BlockType = "fenced_code"
  Html          BlockType = "html"
  Link          BlockType = "link"
  BlockQuote    BlockType = "blockquote"
  List          BlockType = "list"
  ListItem      BlockType = "list_item"
  Image         BlockType = "image"
  LineBreak     BlockType = "line_break"
  BlankLine     BlockType = "blank_line"
)

type Block interface {
  CanConsume(line string) bool
  Consume(line string)
  CanExtend() bool
  String() string
  IsOpen() bool
  Close()
}


