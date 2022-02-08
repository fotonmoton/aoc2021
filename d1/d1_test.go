package d1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD1(t *testing.T) {
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
	actual := D1(input)

	assert.EqualValues(t, 7, actual)
}

func TestD1P2(t *testing.T) {
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

	assert.EqualValues(t, 5, D1P2(input))
}
