package golang

import (
	"testing"
	"github.com/jetuuuu/linter/diapason"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	functions, err := Parse("testdata/simple.go")

	assert.NoError(t, err)
	assert.Equal(t, map[string]diapason.Range{
		"Sum": {Min: 2, Max: 3},
		"IncOne": {Min: 2, Max: 2},
	}, functions)
}

