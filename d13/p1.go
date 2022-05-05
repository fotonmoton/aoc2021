package d13

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type fold struct {
	line int
	axis string
}

type paper map[point]bool

func getPaperAndFolds(in string) (paper, []fold) {
	paper := make(paper, 0)
	folds := make([]fold, 0)

	parts := strings.Split(strings.TrimSpace(in), "\n\n")

	for _, pnt := range strings.Split(parts[0], "\n") {
		xy := strings.Split(pnt, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])

		paper[point{x, y}] = true
	}

	for _, fld := range strings.Split(parts[1], "\n") {
		dirAndLine := strings.Split(strings.Split(fld, " ")[2], "=")
		line, _ := strconv.Atoi(dirAndLine[1])

		folds = append(folds, fold{line, dirAndLine[0]})
	}

	return paper, folds
}

func foldUp(paper paper, y int) paper {
	for pnt := range paper {
		if pnt.y > y {
			newY := y - (pnt.y - y)
			delete(paper, pnt)
			paper[point{pnt.x, newY}] = true
		}
	}

	return paper
}

func foldLeft(paper paper, x int) paper {
	for pnt := range paper {
		if pnt.x > x {
			newX := x - (pnt.x - x)
			delete(paper, pnt)
			paper[point{newX, pnt.y}] = true
		}
	}

	return paper
}

func P1(in string) int {
	paper, folds := getPaperAndFolds(in)

	for _, fold := range folds {
		if fold.axis == "y" {
			paper = foldUp(paper, fold.line)
		} else {
			paper = foldLeft(paper, fold.line)
		}
		break
	}

	return len(paper)
}
