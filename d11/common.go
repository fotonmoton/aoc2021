package d11

import (
	"strconv"
	"strings"
)

type octopus struct {
	energy  int
	flashed bool
}

type grid [][]octopus

func createGrid(in string) grid {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	grid := make(grid, len(lines))

	for i, line := range lines {
		energies := strings.Split(line, "")
		grid[i] = make([]octopus, len(energies))
		for j, energy := range energies {
			num, _ := strconv.Atoi(energy)
			grid[i][j] = octopus{num, false}
		}
	}

	return grid
}

func incr(x, y int, grid grid) int {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return 0
	}

	if grid[y][x].flashed {
		return 0
	}

	grid[y][x].energy++

	if grid[y][x].energy == 10 {
		grid[y][x].energy = 0
		grid[y][x].flashed = true
		return incr(x-1, y-1, grid) +
			incr(x, y-1, grid) +
			incr(x+1, y-1, grid) +
			incr(x-1, y, grid) +
			incr(x+1, y, grid) +
			incr(x-1, y+1, grid) +
			incr(x, y+1, grid) +
			incr(x+1, y+1, grid) + 1
	}

	return 0
}
