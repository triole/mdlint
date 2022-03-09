package main

func main() {
	parseArgs()

	doc := parseMarkdown(CLI.Filename)
	doc.validate()
}
