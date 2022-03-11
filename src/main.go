package main

import (
	"os"
)

func main() {
	parseArgs()
	conf := initConf()

	mdFileList := []string{conf.CLI.Target}
	if isFolder(conf.CLI.Target) == true {
		mdFileList = find(conf.CLI.Target, conf.CLI.Filter)
	}

	for _, mdFile := range mdFileList {
		doc := initDocument(mdFile, conf)
		doc.validate()
		doc.printOutput()
	}
	os.Exit(exitCode)
}
