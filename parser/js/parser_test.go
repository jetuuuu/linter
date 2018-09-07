package js

import (
	"github.com/jetuuuu/linter/diapason"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAndCheck(t *testing.T) {
	gojaFunctions := map[string]diapason.Range{
		"IncOne": {Min: 1, Max: 2},
		"Sum":    {Min: 2, Max: 3},
	}
	problems := ParseAndCheck("testdata", gojaFunctions)

	assert.Equal(t, 2, len(problems))

	mapProblems := make(map[string][]Problem)
	for _, p := range problems {
		mapProblems[p.Function] = append(mapProblems[p.Function], p)
	}

	incOne := Problem{
		File:         "testdata/dir_one/one.js",
		Function:     "IncOne",
		Pos:          Position{Line: 13, Pos: 124},
		ActualArgs:   3,
		ExpectedArgs: gojaFunctions["IncOne"],
	}

	sum := Problem{
		File:         "testdata/simple.js",
		Function:     "Sum",
		Pos:          Position{Line: 15, Pos: 129},
		ActualArgs:   1,
		ExpectedArgs: gojaFunctions["Sum"],
	}

	assert.Equal(t, map[string][]Problem{
		"IncOne": {incOne},
		"Sum":    {sum},
	}, mapProblems)
}
