package core

import (
  "fmt"
  "os"
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
  
  return generateBlockStructure(file) 
}

func generateBlockStructure(file *os.File) (*Document, error) {
  document := NewDocument()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    document.IngestLine(line)
  }

  return &document, nil
}

