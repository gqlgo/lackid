# lackid

[![pkg.go.dev][gopkg-badge]][gopkg]

`lackid` finds a selection for a type which has id field but the selection does not have id.

```graphql
query GetA() {
    a { # want "type A has id field but selection a does not have id field"
	name
    }
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
package main

import (
        "github.com/gqlgo/awesomeanalyzer"            // It is an example. It does not exist actually.
        "github.com/gqlgo/gqlanalysis/multichecker"
        "github.com/gqlgo/lackid"
)

func main() {
        multichecker.Main(
		awesomeanalyzer.Analyzer,
		lackid.Analyzer,
		// You can add more Analyzers!
	)
}
```

`lackid` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/lackid/cmd/lackid@latest
```

The `lackid` command has two flags, `schema` and `query` which will be parsed and analyzed by lackid's Analyzer.

```sh
$ lackid -schema="server/graphql/schema/**/*.graphql" -query="client/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql" and `query` is `query/*/**.graphql`.

`schema` flag accepts URL for a endpoint of GraphQL server.
`lackid` will get schemas by an introspection query via the endpoint.

```sh
$ lackid -schema="https://example.com" -query="client/**/*.graphql"
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/lackid
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/lackid?status.svg
