package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var (
	fol = "../testdata/front_matter"
)

func TestSplitFrontmatter(t *testing.T) {
	// validate front matter
	validateFrontMatter(
		"good.md", true,
		nil, false,
		t,
	)
	validateFrontMatter(
		"broken.md", false,
		nil, false,
		t,
	)

	// evaluate front matter
	validateFrontMatter(
		"good.md", true,
		map[string]string{"title": "string"}, false,
		t,
	)

	// strictly evaluate front matter
	validateFrontMatter(
		"just_title.md", true,
		map[string]string{"title": "string"}, true,
		t,
	)

	validateFrontMatter(
		"title_and_subtitle.md", false,
		map[string]string{"title": "string"}, true,
		t,
	)

	// all supported types
	validateFrontMatter(
		"types.md", true,
		map[string]string{
			"index": "int",
			"m":     "map",
			"no":    "float",
			"tags":  "slice",
			"title": "string",
		}, true,
		t,
	)
}

func validateFrontMatter(target string, expectation bool, fmkeys map[string]string, fmstrict bool, t *testing.T) {
	conf := initTestConf(target, fmkeys, fmstrict)
	for _, mdFile := range conf.FileList {
		doc := initDocument(mdFile, conf)
		doc.validate()
		if doc.IsValid != expectation {
			t.Errorf(
				"fail %q: %s",
				mdFile, doc.Errors,
			)
		}
	}
}

func initTestConf(filter string, fmkeys map[string]string, fmstrict bool) (conf tConf) {
	testFilesFolder, err := filepath.Abs(fol)
	if err != nil {
		fmt.Printf("error reading test files folder %q", err)
		os.Exit(1)
	}
	conf = tConf{
		CLI: tCLI{
			Target:   testFilesFolder,
			Filter:   filter,
			FmKeys:   fmkeys,
			FmStrict: fmstrict,
		},
		FmKeysIterator: makeAlphaIterator(fmkeys),
	}
	conf.FileList = detectFiles(conf.CLI.Target, conf.CLI.Filter)
	return
}
