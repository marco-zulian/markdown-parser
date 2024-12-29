package core

import (
  "regexp"
  "strings"

  "github.com/marco-zulian/markdown-parser/blocks"
);

func Tokenize(s string) []blocks.Token {
  var tokens []blocks.Token

  re := regexp.MustCompile(`^( {0,3})(#{1,6})([ \t]+|$)`)
  thematicBreakRe := regexp.MustCompile(`^((\*{3,})|(-{3,})|(_{3,}))$`)

  if match := re.Find([]byte(s)); match != nil {
    hashRe := regexp.MustCompile(`#{1,6}`)
    endingRe := regexp.MustCompile(` [# ]+$`)
    headingLevel := len(hashRe.Find([]byte(s)))
    
    trimmedString := endingRe.ReplaceAllString(strings.TrimLeft(s, "# "), "")
    tokens = append(tokens, blocks.NewHeaderToken(trimmedString, headingLevel))
  } else if match := thematicBreakRe.Find([]byte(strings.ReplaceAll(s, " ", ""))); match != nil {
    tokens = append(tokens, blocks.NewThematicBreakBlock())
  } else {
    tokens = append(tokens, blocks.NewParagraphBlock(s))
  }

  return tokens
}
