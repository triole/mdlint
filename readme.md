# Mdlint

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
  -h, --help              Show context-sensitive help.
  -f, --filter=<regex>    file detection filter when folder given, default is
                          '\.md$'
  -o, --invalid-only      print out validation result of invalid files only
  -V, --version-flag      display version
```
