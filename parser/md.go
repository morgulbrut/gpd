package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/morgulbrut/gpd/utils"
)

type md struct {
	title  string
	author string
	date   string
	pages  []page
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

func Parse(fileName string) {

	dat, err := ioutil.ReadFile(fileName)
	utils.Check(err)
	input := string(dat)

	Md(input)
}

func Md(doc string) md {
	var d md

	re := regexp.MustCompile("\n---.*\n")
	pgs := re.Split(doc, -1)

	meta := strings.Split(pgs[0], "\n")
	data := pgs[1:]

	for _, l := range meta {
		if strings.HasPrefix(l, "%title") {
			d.title = strings.TrimSpace(strings.Split(l, ":")[1])
		}
		if strings.HasPrefix(l, "%author") {
			d.author = strings.TrimSpace(strings.Split(l, ":")[1])
		}
		if strings.HasPrefix(l, "%date") {
			d.date = strings.TrimSpace(strings.Split(l, ":")[1])
		}
	}
	for _, p := range data {
		Page(p)
	}

	return d
}

func Page(pg string) page {
	var p page
	var e element
	lines := strings.Split(pg, "\n")

	for _, l := range lines {
		if strings.HasPrefix(l, "# ") {
			e.text = strings.Trim(l, "# ")
			e.text = strings.TrimSpace(e.text)
			e.format = Heading1
			p.elements = append(p.elements, e)
		} else if strings.HasPrefix(l, "## ") {
			e.text = strings.Trim(l, "## ")
			e.text = strings.TrimSpace(e.text)
			e.format = Heading2
			p.elements = append(p.elements, e)
		} else if strings.HasPrefix(l, "### ") {
			e.text = strings.Trim(l, "### ")
			e.text = strings.TrimSpace(e.text)
			e.format = Heading3
			p.elements = append(p.elements, e)
		} else {

		}
	}
	for _, bla := range p.elements {
		fmt.Printf("%s: %d\n", bla.text, bla.format)
	}
	return p
}
