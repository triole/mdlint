package main

import "testing"

func TestSplitFrontmatter(t *testing.T) {
	validateSplitFrontMatter("../testdata/good.md", true, t)
	validateSplitFrontMatter("../testdata/broken.md", false, t)
}

func validateSplitFrontMatter(filename string, expectation bool, t *testing.T) {
	doc := parseMarkdown(filename)
	if doc.IsValid != expectation {
		t.Errorf(
			"Split front matter failed: %s, exp: %v, res: %v",
			filename, expectation, doc.IsValid,
		)
	}
}
