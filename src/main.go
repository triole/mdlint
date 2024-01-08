package main

import (
	"fmt"
	"os"
)

func main() {
	parseArgs()
	conf := initConf()

	errorCount := 0
	for _, mdFile := range conf.FileList {
		doc := initDocument(mdFile, conf)
		doc.validate()
		doc.printOutput()
		if !doc.IsValid {
			errorCount++
		}
	}
	if errorCount > 0 {
		fmt.Printf("%v invalid file(s)\n", errorCount)
		os.Exit(1)
	}
}
