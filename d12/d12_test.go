package d12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	in := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	assert.EqualValues(t, 10, P1(in))
}

func TestP2(t *testing.T) {
	in := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	assert.EqualValues(t, 36, P2(in))
}
