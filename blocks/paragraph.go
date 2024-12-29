package blocks

type ParagraphBlock struct {
	tokenType TokenType
	content   string
}

func NewParagraphBlock(content string) ParagraphBlock {
	return ParagraphBlock{
		tokenType: Paragraph,
		content:   content,
	}
}

func (paragraph ParagraphBlock) GetTokenType() TokenType {
	return paragraph.tokenType
}

func (paragraph ParagraphBlock) GetContent() *string {
	return &paragraph.content
}
