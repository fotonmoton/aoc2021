package d13

import "fmt"

func P2(in string) int {
	paper, folds := getPaperAndFolds(in)

	for _, fold := range folds {
		if fold.axis == "y" {
			paper = foldUp(paper, fold.line)
		} else {
			paper = foldLeft(paper, fold.line)
		}
	}

	xMax := 0
	yMax := 0

	for point := range paper {
		if point.x > xMax {
			xMax = point.x
		}

		if point.y > yMax {
			yMax = point.y
		}
	}

	lines := make([][]string, yMax+1)

	for y := range lines {
		lines[y] = make([]string, xMax+1)
	}

	for point := range paper {
		lines[point.y][point.x] = "#"
	}

	for _, line := range lines {
		for _, dot := range line {
			if dot != "" {
				fmt.Print("#")
			} else {
				fmt.Print(" ")

			}
		}
		fmt.Println()
	}

	return len(paper)
}
