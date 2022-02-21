package d1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := `
		199
		200
		208
		210
		200
		207
		240
		269
		260
		263
	`
	actual := P1(input)

	assert.EqualValues(t, 7, actual)
}

func TestP2(t *testing.T) {
	input := `
		199
		200
		208
		210
		200
		207
		240
		269
		260
		263
	`

	assert.EqualValues(t, 5, P2(input))
}
