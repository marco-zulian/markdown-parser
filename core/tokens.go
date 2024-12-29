package core

type TokenType string
const (
  Header TokenType = "header"
)

type Token interface {
  GetTokenType() TokenType
  GetContent() *string 
}

type HeaderToken struct {
  tokenType TokenType
  content   string
  Level     int
}

func NewHeaderToken(content string, level int) HeaderToken {
  return HeaderToken{
    tokenType: Header,
    content: content,
    Level: level,
  }
}

func (header HeaderToken) GetTokenType() TokenType {
  return header.tokenType
}

func (header HeaderToken) GetContent() *string {
  return &header.content
}
