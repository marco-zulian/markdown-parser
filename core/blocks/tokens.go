package core

type BlockType string

const (
	Paragraph BlockType = "paragraph"
	Header    BlockType = "header"
)

type Block interface {
	GetBlockType() BlockType
	GetContent() *string
}
