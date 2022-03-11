# Mdlint ![example workflow](https://github.com/triole/mdlint/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

A simple markdown linter that currently only evaluates if the file's front matter is valid.

## Help

```go mdox-exec="r -h"

simple mdlinter that checks the document's front matter

Arguments:
  [<target>]    can be file or folder, if folder files to process are detected
                recursively

Flags:
  -h, --help                    Show context-sensitive help.
  -f, --filter=REGEX            file detection filter when folder given, default
                                is '\.md$'
  -e, --fmkeys=KEY=VALUE;...    evaluate a distinct front matter key and its
                                value type, can be used multiple times (i.e. -e
                                index=int -e no=float -e title=string -e
                                tags=slice -e m=map)
  -s, --fmstrinct               strictly evaluate front matter, documents are
                                considered to be invalid if front matter
                                contains a key that wasn't provided with the
                                'fmkeys' arg
  -o, --invalid-only            print out validation result of invalid files
                                only
  -V, --version-flag            display version
```
