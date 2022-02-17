package d6

import (
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

func D6(input string, startDays int) int {
	fishes := toNum(strings.Split(strings.TrimSpace(input), ","))

	for day := 0; day < startDays; day++ {
		for idx := range fishes {
			if fishes[idx] == 0 {
				fishes[idx] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[idx]--
			}
		}
	}

	return len(fishes)
}

// Inspired by
// https://www.reddit.com/r/adventofcode/comments/r9z49j/comment/hnfaisu/?utm_source=share&utm_medium=web2x&context=3
func D6P2(input string, startDays int) int {
	fishes := toNum(strings.Split(strings.TrimSpace(input), ","))
	ages := make([]int, 9)

	for _, fish := range fishes {
		ages[fish]++
	}

	for day := 0; day < startDays; day++ {
		dieToday := ages[0]
		ages = ages[1:]
		ages[6] += dieToday
		ages = append(ages, dieToday)
	}

	sum := 0
	for _, age := range ages {
		sum += age
	}

	return sum
}
