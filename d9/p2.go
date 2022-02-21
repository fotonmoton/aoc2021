package d9

import (
	"sort"
	"strings"
)

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

func P2(in string) int {
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
