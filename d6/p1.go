package d6

import (
	"strings"
)

func P1(input string, startDays int) int {
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
