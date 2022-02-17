package d7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func toNum(in []string) []int {
	nums := make([]int, len(in))

	for idx, str := range in {
		num, _ := strconv.Atoi(str)
		nums[idx] = num
	}

	return nums
}

func D7(in string) int {
	positions := toNum(strings.Split(strings.TrimSpace(in), ","))

	sort.Sort(sort.IntSlice(positions))

	nearest := positions[int(math.Round(float64(len(positions)/2)))]

	fuel := 0
	for _, crab := range positions {
		fuel += int(math.Abs(float64(crab - nearest)))
	}

	return fuel
}

func sumUp(to int) int {
	result := 0

	for from := 0; from <= to; from++ {
		result += from
	}

	return result
}

func D7P2(in string) int {
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
