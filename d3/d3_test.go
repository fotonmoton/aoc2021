package d3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD3(t *testing.T) {
	input := `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	assert.EqualValues(t, 198, D3(input))
}

func TestD3P2(t *testing.T) {
	input := `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	assert.EqualValues(t, 230, D3P2(input))
}
