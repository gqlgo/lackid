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
