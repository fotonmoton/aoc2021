package d11

func isAllFlashed(grid grid) bool {

	for y, line := range grid {
		for x := range line {
			if !grid[y][x].flashed {
				return false
			}
		}
	}

	return true
}

func flash2(grid grid) int {
	step := 0
	for {
		for y, line := range grid {
			for x := range line {
				incr(x, y, grid)
			}
		}
		step++
		if isAllFlashed(grid) {
			return step
		}
		reset(grid)
	}
}

func P2(in string) int {
	grid := createGrid(in)
	return flash2(grid)
}
