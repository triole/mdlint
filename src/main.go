package main

import (
	"os"
)

func main() {
	parseArgs()
	conf := initConf()

	for _, mdFile := range conf.FileList {
		doc := initDocument(mdFile, conf)
		doc.validate()
		doc.printOutput()
	}
	os.Exit(exitCode)
}
