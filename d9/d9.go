package d9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func bigger(a byte, b byte) bool {
	aNum, _ := strconv.Atoi(string(a))
	bNum, _ := strconv.Atoi(string(b))

	return aNum > bNum
}

func toNum(str byte) int {
	num, _ := strconv.Atoi(string(str))
	return num
}

func addWalls(lines []string) []string {
	for i, line := range lines {
		lines[i] = fmt.Sprintf("9%v9", line)
	}

	nines := strings.Repeat("9", len(lines[0]))
	return append([]string{nines}, append(lines, nines)...)
}

func D9(in string) int {
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

type exists struct{}

type pos struct {
	x int
	y int
}

type basin map[pos]exists

const wall = '9'

func flooded(x, y int, basins []basin) bool {
	for _, basin := range basins {
		if _, ok := basin[pos{x, y}]; ok {
			return true
		}
	}

	return false
}

func flood(x, y int, heights []string, basin basin) basin {
	if heights[y][x] == wall {
		return basin
	}

	if _, ok := basin[pos{x, y}]; ok {
		return basin
	}

	basin[pos{x, y}] = exists{}

	flood(x, y-1, heights, basin)
	flood(x+1, y, heights, basin)
	flood(x, y+1, heights, basin)
	flood(x-1, y, heights, basin)
	return basin
}

func D9P2(in string) int {
	heights := addWalls(strings.Split(strings.TrimSpace(in), "\n"))
	basins := []basin{}

	for y := 1; y < len(heights)-1; y++ {
		for x := 1; x < len(heights[0])-1; x++ {
			if !flooded(x, y, basins) && heights[y][x] != wall {
				basins = append(basins, flood(x, y, heights, basin{}))
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool { return len(basins[i]) > len(basins[j]) })

	return len(basins[0]) * len(basins[1]) * len(basins[2])
}
