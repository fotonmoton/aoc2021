package d9

import (
	"strconv"
	"strings"
)

func toNum(str byte) int {
	num, _ := strconv.Atoi(string(str))
	return num
}

func bigger(a byte, b byte) bool {
	return toNum(a) > toNum(b)
}

func P1(in string) int {
	lines := addWalls(strings.Split(strings.TrimSpace(in), "\n"))
	risk := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			current := lines[i][j]
			top := bigger(lines[i-1][j], current)
			right := bigger(lines[i][j+1], current)
			bottom := bigger(lines[i+1][j], current)
			left := bigger(lines[i][j-1], current)
			if top && right && bottom && left {
				risk += toNum(current) + 1
			}
		}
	}

	return risk
}
