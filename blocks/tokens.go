package blocks

type BlockType string

const (
	Paragraph     BlockType = "paragraph"
	Header        BlockType = "header"
	ThematicBreak BlockType = "break"
	IndentedCode  BlockType = "indented_code"
)

type Block interface {
	GetBlockType() BlockType
	GetContent() *string
}
