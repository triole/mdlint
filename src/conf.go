package main

type tConf struct {
	CLI            tCLI
	FmKeysIterator []string
	FileList       []string
}

type tCLI struct {
	Target      string
	Filter      string
	FmKeys      map[string]string
	FmStrict    bool
	InvalidOnly bool
}

func initConf() (conf tConf) {
	conf = tConf{
		CLI: tCLI{
			Target:      CLI.Target,
			Filter:      CLI.Filter,
			FmKeys:      CLI.Fmkeys,
			FmStrict:    CLI.Fmstrinct,
			InvalidOnly: CLI.InvalidOnly,
		},
		FmKeysIterator: makeAlphaIterator(CLI.Fmkeys),
	}
	conf.FileList = detectFiles(conf.CLI.Target, conf.CLI.Filter)
	return
}

func detectFiles(target, filter string) (fileList []string) {
	fileList = []string{target}
	if isFolder(target) == true {
		fileList = find(target, filter)
	}
	return
}
