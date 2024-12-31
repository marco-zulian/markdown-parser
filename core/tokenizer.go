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

//func Tokenize(s string) []blocks.Block {
//	if match := codeRe.Find([]byte(s)); match != nil {
//    result = append(result, blocks.NewCodeBlock(s[TAB_SIZE:])) // TODO Still placing paragraph
//	} else if match := headingRe.Find([]byte(s)); match != nil {
//		result = append(result, blocks.NewHeaderBlock(trimmedString, headingLevel))
//	} else if match := thematicBreakRe.Find([]byte(strings.ReplaceAll(s, " ", ""))); match != nil {
//		result = append(result, blocks.NewThematicBreakBlock())
//	} else {
//		result = append(result, blocks.NewParagraphBlock(s))
//	}
//
//	return result 
//}

