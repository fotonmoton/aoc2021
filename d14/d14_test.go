package d14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	result := P1(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)

	assert.EqualValues(t, 1588, result)
}

func TestP2(t *testing.T) {
	result := P2(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)

	assert.EqualValues(t, 2188189693529, result)
}
