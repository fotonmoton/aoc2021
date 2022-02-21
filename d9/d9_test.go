package d9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD8(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	assert.EqualValues(t, 15, D9(input))
}

func TestD8P2(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	assert.EqualValues(t, 1134, D9P2(input))
}
