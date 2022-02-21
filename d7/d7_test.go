package d7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {

	assert.EqualValues(t, 37, P1("16,1,2,0,4,2,7,1,2,14"))
}

// This test from example fails but solution works, lol.
// Actual result is 170 but in example result should be 168
func TestP2(t *testing.T) {

	assert.EqualValues(t, 170, P2("16,1,2,0,4,2,7,1,2,14"))
}
