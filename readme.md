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

Flags:
  -h, --help            Show context-sensitive help.
  -e, --errors-only     print only files where errors occured
  -V, --version-flag    display version
```
