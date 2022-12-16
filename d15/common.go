package d15

import (
	"strconv"
	"strings"
)

func weightNeighbors(in string) (weights, neighbors) {
	weights := make(weights)
	neighbors := make(neighbors)
	lines := strings.Split(strings.TrimSpace(in), "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, weight := range strings.Split(line, "") {
			weight, _ := strconv.Atoi(weight)
			weights[node(x+y*height)] = weight
		}
	}
	maxX := width
	maxY := height
	for height >= 0 {
		width = maxX
		for width >= 0 {
			current := node(width + height*maxY)
			if height > 0 {
				neighbors[current] = append(neighbors[current], node(width+(height-1)*maxY))
			}
			if height < maxY {
				neighbors[current] = append(neighbors[current], node(width+(height+1)*maxY))
			}
			if width > 0 {
				neighbors[current] = append(neighbors[current], node(width+height*maxY-1))
			}
			if width < maxX {
				neighbors[current] = append(neighbors[current], node(width+height*maxY+1))

			}
			width--
		}
		height--
	}
	return weights, neighbors
}
