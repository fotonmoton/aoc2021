package d1

import (
	"strconv"
	"strings"
)

func D1(scans string) int64 {
	depths := strings.Split(strings.TrimSpace(scans), "\n")
	downs := 0

	for i := 1; i < len(depths); i++ {
		next, _ := strconv.Atoi(strings.TrimSpace(depths[i]))
		prev, _ := strconv.Atoi(strings.TrimSpace(depths[i-1]))

		if next > prev {
			downs++
		}
	}

	return int64(downs)
}

func D1P2(scans string) int64 {
	depths := strings.Split(strings.TrimSpace(scans), "\n")
	downs := 0

	for i := 3; i < len(depths); i++ {
		fourth, _ := strconv.Atoi(strings.TrimSpace(depths[i]))
		third, _ := strconv.Atoi(strings.TrimSpace(depths[i-1]))
		second, _ := strconv.Atoi(strings.TrimSpace(depths[i-2]))
		first, _ := strconv.Atoi(strings.TrimSpace(depths[i-3]))

		prev := first + second + third
		next := second + third + fourth

		if next > prev {
			downs++
		}
	}

	return int64(downs)
}
