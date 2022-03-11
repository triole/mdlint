package main

import "testing"

func TestSplitFrontmatter(t *testing.T) {
	validateSplitFrontMatter("../testdata/good.md", true, t)
	validateSplitFrontMatter("../testdata/broken.md", false, t)
}

func validateSplitFrontMatter(filename string, expectation bool, t *testing.T) {
	conf := tConf{}
	doc := initDocument(filename, conf)
	doc.validate()
	if doc.IsValid != expectation {
		t.Errorf(
			"Split front matter failed: %s, exp: %v, res: %v",
			filename, expectation, doc.IsValid,
		)
	}
}
