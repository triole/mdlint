package main

import "os"

func main() {
	parseArgs()

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
