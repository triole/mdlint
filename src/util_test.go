package main

import "testing"

func TestIsFolder(t *testing.T) {
	validateIsFolder("../testdata", true, t)
	validateIsFolder("../go.mod", false, t)
}

func validateIsFolder(target string, exp bool, t *testing.T) {
	res := isFolder(target)
	if res != exp {
		t.Errorf(
			"isfolder failed: %q returned %v not %v",
			target, res, exp,
		)
	}
}

func TestRxFind(t *testing.T) {
	validateRxFind(`^abc`, "abcdefg", "abc", t)
	validateRxFind(`efg$`, "abcdefg", "efg", t)
}

func validateRxFind(rx, str, exp string, t *testing.T) {
	res := rxFind(rx, str)
	if res != exp {
		t.Errorf(
			"rxfind failed: %q on %q delivered %q not %q",
			rx, str, res, exp,
		)
	}
}
