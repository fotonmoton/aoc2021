package d7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD7(t *testing.T) {

	assert.EqualValues(t, 37, D7("16,1,2,0,4,2,7,1,2,14"))
}

// This test from example fails but solution works, lol
func TestD7P2(t *testing.T) {

	assert.EqualValues(t, 168, D7P2("16,1,2,0,4,2,7,1,2,14"))
}
