package core

import (
  "fmt"
  "os"
  "io"
  "bufio"
)

const TAB_SIZE = 4

func Tokenize(filePath string) (*Document, error) {
  file, err := os.Open(filePath)
  defer file.Close()

  if err != nil {
    fmt.Errorf("Could not load file at path %s. %q", filePath, err) 
    return nil, err
  }
  
  return GenerateBlockStructure(file) 
}

func GenerateBlockStructure(inputStream io.Reader) (*Document, error) {
  document := NewDocument()

  scanner := bufio.NewScanner(inputStream)
  for scanner.Scan() {
    line := scanner.Text()
    document.IngestLine(line)
  }

  return &document, nil
}

