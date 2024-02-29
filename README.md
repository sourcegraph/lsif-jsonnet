# JSonnet LSIF indexer ![](https://img.shields.io/badge/status-development-yellow)

🚨 This implementation is still in very early stage and follows the latest LSIF specification closely.

## Language Server Index Format

The purpose of the Language Server Index Format (LSIF) is to define a standard format for language servers or other
programming tools to dump their knowledge about a workspace. This dump can later be used to answer language server
[LSP](https://microsoft.github.io/language-server-protocol/) requests for the same workspace without running the
language server itself. Since much of the information would be invalidated by a change to the workspace,
the dumped information typically excludes requests used when mutating a document. So, for example, the result of a code
complete request is typically not part of such a dump.

A first draft specification can be found [here](https://github.com/Microsoft/language-server-protocol/blob/master/indexFormat/specification.md).

## JSonnet

[JSonnet](https://jsonnet.org) is a data templating language related to JSON.

## Implementation

In true open-source spirit the implementation is made from these parts:
 
- The parser is generated by [Antlr4](https://www.antlr.org) from a grammar file modified from
 [this GitHub Gist](https://gist.github.com/ironchefpython/84380aa60871853dc86719dd598c35e4).
 
- The protocol.go code was borrowed and modified from [lsif-go](https://github.com/sourcegraph/lsif-go).

- The LSIF dumper is a modification of indexer.go  from [lsif-go](https://github.com/sourcegraph/lsif-go).

- The scope implementation is a modification of [go.types.Scope](https://golang.org/pkg/go/types/#Scope).

Many thanks to the creators of these.

## Status

This was put together quickly to learn LSIF and JSonnet. It is missing many features, for example
hover results, documentation comments (parser currently discards comments), error handling, the standard library is not covered.
It has not seen code reviews, so it's highly likely that there are bugs, omissions and oversights. 
It can serve as a starting point for more production-ready implementations for JSonnet or other languages.

## Usage

- Build the `cmd/lsif-jsonnet` cli the usual way with Go.
- The cli mimics the flags of the `jsonnet` cli, so you can specify `libsonnet` search paths with `-J` and the output
file with `-o`.
- Run the cli on your jsonnet file and generate an LSIF data file.
- You can use the generated LSIF data file in [VSCode](https://code.visualstudio.com) with the LSIF plugin or when running a
  [Sourcegraph](https://docs.sourcegraph.com/user/code_intelligence/lsif) instance. Feel free to fork and use the
  [test-jsonnet-lsif](https://github.com/uwedeportivo/test-jsonnet-lsif) repo as your test code.
  
## Re-generate parser

After installing [Antlr4](https://github.com/antlr/antlr4) and the [Go target](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)
you can run this command from the repo root:

```bash
antlr -Dlanguage=Go -o parser Jsonnet.g4
```

## Testing and validating LSIF output

[lsif-util](https://www.npmjs.com/package/lsif-util) is a commandline tool for validating the generated LSIF data syntactically.
It also supports producing [dot graphs](https://graphviz.gitlab.io/_pages/doc/info/lang.html) that can be turned into PNG images
with [dot](https://graphviz.gitlab.io/download/).

- Validate: `lsif-util validate data.lsif`
- Visualize: `lsif-util visualize data.lsif --distance 2 | dot -Tpng -o image.png`
