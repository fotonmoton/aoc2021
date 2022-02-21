package d8

import (
	"strings"
)

func P1(in string) int {
	lines := strings.Split(strings.TrimSpace(in), "\n")

	sum := 0

	for _, line := range lines {
		for _, digit := range strings.Split(strings.Split(line, "|")[1], " ") {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}
	}

	return sum
}
