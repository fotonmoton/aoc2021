package d7

import (
	"math"
	"sort"
	"strings"
)

func P1(in string) int {
	positions := toNum(strings.Split(strings.TrimSpace(in), ","))

	sort.Sort(sort.IntSlice(positions))

	nearest := positions[int(math.Round(float64(len(positions)/2)))]

	fuel := 0
	for _, crab := range positions {
		fuel += int(math.Abs(float64(crab - nearest)))
	}

	return fuel
}
