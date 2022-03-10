package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adrg/frontmatter"
)

type tMatter map[string]interface{}
type tDocument struct {
	Filename    string
	FullContent tStringByte
	FrontMatter tMatter
	Rest        tStringByte
	Errors      []error
	IsValid     bool
}

type tStringByte struct {
	Bytes  []byte
	String string
}

func parseMarkdown(filename string) (doc tDocument) {
	doc.Filename = filename
	doc.readFile()
	doc.splitFrontMatter()
	return doc
}

func (doc *tDocument) readFile() {
	var err error
	doc.FullContent.Bytes, err = ioutil.ReadFile(doc.Filename)
	doc.addError(err)
	doc.FullContent.String = string(doc.FullContent.Bytes)
	return
}

func (doc *tDocument) splitFrontMatter() {
	var mat tMatter
	var err error
	doc.Rest.Bytes, err = frontmatter.Parse(
		strings.NewReader(doc.FullContent.String), &mat,
	)
	doc.addError(err)
	doc.FrontMatter = mat
	doc.Rest.String = string(doc.Rest.Bytes)
	return
}

func (doc *tDocument) validate() {
	if len(doc.Errors) > 0 {
		fmt.Printf("%-7s %q, %s\n", "Invalid", doc.Filename, doc.Errors)
		exitCode = 1
	} else {
		if CLI.InvalidOnly == false {
			fmt.Printf("%-7s %q\n", "Ok", doc.Filename)
		}
	}
}

func (doc *tDocument) addError(err error) {
	if err != nil {
		doc.Errors = append(doc.Errors, err)
	}
	doc.IsValid = doc.isValid()
}

func (doc *tDocument) isValid() (b bool) {
	return len(doc.Errors) == 0
}
