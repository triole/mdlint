package main

type tConf struct {
	CLI            tCLI
	FmKeysIterator []string
}

type tCLI struct {
	Target      string
	Filter      string
	FmKeys      map[string]string
	FmStrict    bool
	InvalidOnly bool
}

func initConf() (conf tConf) {
	return tConf{
		CLI: tCLI{
			Target:      CLI.Target,
			Filter:      CLI.Filter,
			FmKeys:      CLI.Fmkeys,
			FmStrict:    CLI.Fmstrinct,
			InvalidOnly: CLI.InvalidOnly,
		},
		FmKeysIterator: makeAlphaIterator(CLI.Fmkeys),
	}
}
