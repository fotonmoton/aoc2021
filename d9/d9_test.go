package d9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	assert.EqualValues(t, 15, P1(input))
}

func TestP2(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	assert.EqualValues(t, 1134, P2(input))
}
