package parser

import (
	"fmt"

	mdParser "github.com/nikitavoloboev/markdown-parser"
)

type md struct {
	title string
	autor string
	data  string
	pages []page
}

type page struct {
	elements []element
}

type element struct {
	text   string
	format format
}

type format int

const (
	Default format = iota
	Bold
	Italic
	BoldItalic
	Heading1
	Heading2
	Heading3
	Quote
	Code
)

func Md(fileName string) {
	// Open the file
	mdFile, _ := mdParser.ParseMarkdownFile(fileName)

	fmt.Println(mdFile)
}
