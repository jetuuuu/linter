package diapason

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange_Contains(t *testing.T) {
	r := Range{Min: 100, Max: 500}

	assert.True(t, r.Contains(200))
	assert.False(t, r.Contains(0))

}

func TestRange_String(t *testing.T) {
	assert.Equal(t, "between 100 and 500", Range{100, 500}.String())
	assert.Equal(t, "40", Range{Min: 40, Max: 40}.String())
}
