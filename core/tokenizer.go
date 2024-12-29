package core

import (
  "regexp"
  "strings"

  "github.com/marco-zulian/markdown-parser/blocks"
);

func Tokenize(s string) []blocks.Token {
  var tokens []blocks.Token

  re := regexp.MustCompile(`^( {0,3})(#{1,6})([ \t]+|$)`)
  if match := re.Find([]byte(s)); match != nil {
    hashRe := regexp.MustCompile(`#{1,6}`)
    endingRe := regexp.MustCompile(` [# ]+$`)
    headingLevel := len(hashRe.Find([]byte(s)))
    
    trimmedString := endingRe.ReplaceAllString(strings.TrimLeft(s, "# "), "")
    tokens = append(tokens, blocks.NewHeaderToken(trimmedString, headingLevel))
  } else {
    tokens = append(tokens, blocks.NewParagraphBlock(s))
  }

  return tokens
}
