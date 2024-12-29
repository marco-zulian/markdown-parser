package core

import (
  "strings"
);

func Tokenize(s string) []Token {
  var tokens []Token

  if (strings.HasPrefix(s, "#")) {
    tokens = append(tokens, NewHeaderToken(s[1:], 1))
  }

  return tokens
}
