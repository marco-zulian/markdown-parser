package core

import (
	"regexp"
	"strings"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func Blockize(s string) []blocks.Block {
	var Blocks []blocks.Block

	re := regexp.MustCompile(`^( {0,3})(#{1,6})([ \t]+|$)`)
	indentedCodeRe := regexp.MustCompile(`^ {4,}`)
	thematicBreakRe := regexp.MustCompile(`^((\*{3,})|(-{3,})|(_{3,}))$`)

	if match := indentedCodeRe.Find([]byte(s)); match != nil {
		Blocks = append(Blocks, blocks.NewParagraphBlock(s)) // TODO Still placing paragraph
	} else if match := re.Find([]byte(s)); match != nil {
		hashRe := regexp.MustCompile(`#{1,6}`)
		endingRe := regexp.MustCompile(` [# ]+$`)
		headingLevel := len(hashRe.Find([]byte(s)))

		trimmedString := endingRe.ReplaceAllString(strings.TrimLeft(s, "# "), "")
		Blocks = append(Blocks, blocks.NewHeaderBlock(trimmedString, headingLevel))
	} else if match := thematicBreakRe.Find([]byte(strings.ReplaceAll(s, " ", ""))); match != nil {
		Blocks = append(Blocks, blocks.NewThematicBreakBlock())
	} else {
		Blocks = append(Blocks, blocks.NewParagraphBlock(s))
	}

	return Blocks
}
