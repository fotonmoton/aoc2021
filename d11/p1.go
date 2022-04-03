package d11

func reset(grid grid) {
	for y, line := range grid {
		for x := range line {
			grid[y][x].flashed = false
		}
	}
}

func flash(grid grid) int {
	lights := 0
	for i := 0; i < 100; i++ {
		for y, line := range grid {
			for x := range line {
				lights += incr(x, y, grid)
			}
		}
		reset(grid)
	}

	return lights
}

func P1(in string) int {
	grid := createGrid(in)
	return flash(grid)
}
