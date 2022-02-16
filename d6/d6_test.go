package d6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD6(t *testing.T) {
	assert.EqualValues(t, 5934, D6("3,4,3,1,2", 80))
}

func TestD6P2Naive(t *testing.T) {
	assert.EqualValues(t, 26984457539, D6("3", 256))
}

func TestD6P2(t *testing.T) {
	assert.EqualValues(t, 26984457539, D6P2("3,4,3,1,2", 256))
}
