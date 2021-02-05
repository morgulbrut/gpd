package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/morgulbrut/gpd/utils"
)

type Md struct {
	Title  string
	Author string
	Date   string
	Pages  []Page
}

type Page struct {
	Elements []Element
}

type Element struct {
	Text   string
	Format Format
}

type Format int

const (
	DEFAULT Format = iota
	BOLD
	ITALIC
	BOLDITALIC
	HEADING1
	HEADING2
	HEADING3
	HEADING4
	HEADING5
	HEADING6
	BLOCKQOUTE
	ULIST
	OLIST
	CODE
)

func Parse(fileName string) {

	dat, err := ioutil.ReadFile(fileName)
	utils.Check(err)
	input := string(dat)

	md(input)
}

func md(doc string) Md {
	var d Md

	re := regexp.MustCompile("\n---.*\n")
	pgs := re.Split(doc, -1)

	meta := strings.Split(pgs[0], "\n")
	data := pgs[1:]

	for _, l := range meta {
		if strings.HasPrefix(l, "%title") {
			d.Title = strings.TrimSpace(strings.Split(l, ":")[1])
		}
		if strings.HasPrefix(l, "%author") {
			d.Author = strings.TrimSpace(strings.Split(l, ":")[1])
		}
		if strings.HasPrefix(l, "%date") {
			d.Date = strings.TrimSpace(strings.Split(l, ":")[1])
		}
	}
	for _, p := range data {
		page(p)
	}

	return d
}

func page(pg string) Page {
	var p Page
	var e Element
	lines := strings.Split(pg, "\n")

	for _, l := range lines {

		//BlockQuotes
		if strings.HasPrefix(l, "> ") {
			e.Text = strings.Trim(l, "> ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = BLOCKQOUTE
			p.Elements = append(p.Elements, e)
		}
		// Headings
		if strings.HasPrefix(l, "# ") {
			e.Text = strings.Trim(l, "# ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = HEADING1
			p.Elements = append(p.Elements, e)
		} else if strings.HasPrefix(l, "## ") {
			e.Text = strings.Trim(l, "## ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = HEADING2
			p.Elements = append(p.Elements, e)
		} else if strings.HasPrefix(l, "### ") {
			e.Text = strings.Trim(l, "### ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = HEADING3
			p.Elements = append(p.Elements, e)
		} else if strings.HasPrefix(l, "#### ") {
			e.Text = strings.Trim(l, "#### ")
			e.Text = strings.TrimSpace(e.Text)i
			e.Format = HEADING4
			p.Elements = append(p.Elements, e)
		} else if strings.HasPrefix(l, "##### ") {
			e.Text = strings.Trim(l, "##### ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = HEADING5
			p.Elements = append(p.Elements, e)
		} else if strings.HasPrefix(l, "###### ") {
			e.Text = strings.Trim(l, "###### ")
			e.Text = strings.TrimSpace(e.Text)
			e.Format = HEADING6
			p.Elements = append(p.Elements, e)

		} else {

		}
	}
	for _, bla := range p.Elements {
		fmt.Printf("%s: %d\n", bla.Text, bla.Format)
	}
	return p
}
