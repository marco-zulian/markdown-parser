package blocks

type ParagraphBlock struct {
	content string
}

func NewParagraphBlock(content string) ParagraphBlock {
	return ParagraphBlock{
		content: content,
	}
}

func (paragraph ParagraphBlock) GetBlockType() BlockType {
	return Paragraph
}

func (paragraph ParagraphBlock) GetContent() *string {
	return &paragraph.content
}
