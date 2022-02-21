package d7

import (
	"math"
	"sort"
	"strings"
)

func P2(in string) int {
	positions := toNum(strings.Split(strings.TrimSpace(in), ","))

	sort.Sort(sort.IntSlice(positions))

	fuelCosts := make([]int, len(positions))

	for idx, crab := range positions {
		costForPosition := 0
		for _, nextCrab := range positions {
			costForPosition += sumUp(int(math.Abs(float64(crab - nextCrab))))
		}
		fuelCosts[idx] = costForPosition
	}

	sort.Sort(sort.IntSlice(fuelCosts))

	return fuelCosts[0]
}
