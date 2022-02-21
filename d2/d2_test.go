package d2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {

	input := `
forward 5
down 5
forward 8
up 3
down 8
forward 2`

	assert.EqualValues(t, 150, P1(input))
}

func TestP2(t *testing.T) {

	input := `
forward 5
down 5
forward 8
up 3
down 8
forward 2`

	assert.EqualValues(t, 900, P2(input))
}
