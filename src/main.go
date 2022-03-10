package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	parseArgs()

	fmt.Printf("%q\n", reflect.ValueOf(map[string]string{}).Kind())

	mdFileList := []string{CLI.Target}
	if isFolder(CLI.Target) == true {
		mdFileList = find(CLI.Target, CLI.Filter)
	}

	for _, mdFile := range mdFileList {
		doc := parseMarkdown(mdFile)
		doc.validate()
	}
	os.Exit(exitCode)
}
