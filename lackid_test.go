package lackid_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/gqlgo/lackid"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, lackid.Analyzer, "a")
}
