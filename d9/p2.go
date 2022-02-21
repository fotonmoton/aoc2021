package d9

import (
	"sort"
	"strings"
)

type pos struct {
	x int
	y int
}

type basin map[pos]bool

func flood(x, y int, heights []string, basin basin, flooded basin) basin {
	if heights[y][x] == wall {
		return basin
	}

	if flooded[pos{x, y}] {
		return basin
	}

	basin[pos{x, y}] = true
	flooded[pos{x, y}] = true

	flood(x, y-1, heights, basin, flooded)
	flood(x+1, y, heights, basin, flooded)
	flood(x, y+1, heights, basin, flooded)
	flood(x-1, y, heights, basin, flooded)

	return basin
}

func P2(in string) int {
	heights := addWalls(strings.Split(strings.TrimSpace(in), "\n"))
	basins := []basin{}
	flooded := basin{}

	for y := 1; y < len(heights)-1; y++ {
		for x := 1; x < len(heights[0])-1; x++ {
			basins = append(basins, flood(x, y, heights, basin{}, flooded))
		}
	}

	sort.Slice(basins, func(i, j int) bool { return len(basins[i]) > len(basins[j]) })

	return len(basins[0]) * len(basins[1]) * len(basins[2])
}
