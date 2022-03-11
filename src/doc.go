package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/adrg/frontmatter"
)

type tMatter map[string]interface{}
type tDocument struct {
	Filename    string
	FullContent tStringByte
	FrontMatter tMatter
	Rest        tStringByte
	Errors      []string
	IsValid     bool
	Conf        tConf
}

type tStringByte struct {
	Bytes  []byte
	String string
}

func initDocument(filename string, conf tConf) (doc tDocument) {
	doc = tDocument{
		Filename: filename,
		Conf:     conf,
	}
	doc.readFile()
	return doc
}

func (doc *tDocument) readFile() {
	var err error
	doc.FullContent.Bytes, err = ioutil.ReadFile(doc.Filename)
	doc.addError(err)
	doc.FullContent.String = string(doc.FullContent.Bytes)
	return
}

func (doc *tDocument) validate() {
	doc.splitFrontMatter()
	if doc.IsValid == true && len(doc.Conf.CLI.FmKeys) > 0 {
		doc.evaluateFrontMatter()
	}
	if doc.IsValid && doc.Conf.CLI.FmStrict == true {
		doc.strictlyEvaluateFrontMatter()
	}
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

func (doc *tDocument) evaluateFrontMatter() {
	for _, key := range doc.Conf.FmKeysIterator {
		val := doc.Conf.CLI.FmKeys[key]
		fmVal := doc.FrontMatter[key]
		if fmVal == nil {
			doc.addError(
				fmt.Errorf(
					"front matter entry %q does not exist", key,
				),
			)
		} else {
			fmValKind := rxFind(
				"^[a-z]+", fmt.Sprintf("%s", reflect.ValueOf(fmVal).Kind()),
			)
			if val != fmValKind {
				fmt.Printf("%q\n", fmVal)
				doc.addError(
					fmt.Errorf(
						"front matter entry %q is %s not %s", key, fmValKind, val,
					),
				)
			}
		}
	}
}

func (doc *tDocument) strictlyEvaluateFrontMatter() {
	iterator := makeAlphaIteratorItf(doc.FrontMatter)
	if reflect.DeepEqual(iterator, doc.Conf.FmKeysIterator) == false {
		doc.addError(
			fmt.Errorf(
				"strictly evaluate front matter failed: its keys are %s not %s",
				iterator, doc.Conf.FmKeysIterator,
			),
		)
	}
}

func (doc *tDocument) printOutput() {
	if len(doc.Errors) > 0 {
		fmt.Printf(
			"%-7s %q, [%s]\n", "Invalid",
			doc.Filename, strings.Join(doc.Errors, ", "),
		)
	} else {
		if doc.Conf.CLI.InvalidOnly == false {
			fmt.Printf("%-7s %q\n", "Ok", doc.Filename)
		}
	}
}

func (doc *tDocument) addError(err error) {
	if err != nil {
		doc.Errors = append(doc.Errors, err.Error())
	}
	doc.IsValid = doc.isValid()
}

func (doc *tDocument) isValid() (b bool) {
	return len(doc.Errors) == 0
}
