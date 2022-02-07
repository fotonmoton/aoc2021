package d2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD2(t *testing.T) {

	input := `
forward 5
down 5
forward 8
up 3
down 8
forward 2`

	assert.EqualValues(t, 150, D2(input))
}
