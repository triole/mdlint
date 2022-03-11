package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "mdlint"
	appDescription = "simple mdlinter that checks the document's front matter"
	appMainversion = "0.1"
)

var CLI struct {
	Target      string            `help:"can be file or folder, if folder files to process are detected recursively" arg optional`
	Filter      string            `help:"file detection filter when folder given, default is '\.md$'" short:f default:\.md$ placeholder:"REGEX"`
	Fmkeys      map[string]string `help:"evaluate a distinct front matter key and its value type, can be used multiple times  (i.e. -e index=int -e no=float -e title=string -e tags=slice -e m=map)" short:e sep:","`
	Fmstrinct   bool              `help:"strictly evaluate front matter, documents are considered to be invalid if front matter contains a key that wasn't provided with the 'fmkeys' arg" short:s`
	InvalidOnly bool              `help:"print out validation result of invalid files only" short:o`
	VersionFlag bool              `help:"display version" short:V`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
	)
	_ = ctx.Run()

	if CLI.VersionFlag == true {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	if CLI.Target == "" {
		fmt.Printf("%s\n", "Error: Positional arg expected. Please pass file or folder name.")
		os.Exit(1)
	}
	// ctx.FatalIfErrorf(err)
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}
