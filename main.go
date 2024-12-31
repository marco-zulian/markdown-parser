package main

import (
  "fmt"

  "github.com/marco-zulian/markdown-parser/core"
)

func main() {
  document, err := core.Tokenize("example.md") 

  if err != nil {
    fmt.Printf("Error tokenizing %q", err)
  }

  fmt.Printf("%s", *document.GetContent())
}
