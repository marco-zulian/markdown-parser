package blocks

type HeaderBlock struct {
	content string
	Level   int
}

func NewHeaderBlock(content string, level int) HeaderBlock {
	return HeaderBlock{
		content: content,
		Level:   level,
	}
}

func (header HeaderBlock) GetBlockType() BlockType {
	return Header
}

func (header HeaderBlock) GetContent() *string {
	return &header.content
}
