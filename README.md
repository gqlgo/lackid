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

```go
package main

import (
        "github.com/gqlgo/gqlanalysis/multichecker"
        "github.com/gqlgo/lackid"
)

func main() {
        multichecker.Main(
		lackid.Analyzer,
	)
}
```

`lackid` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/lackid/cmd/lackid@latest
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/lackid
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/lackid?status.svg
