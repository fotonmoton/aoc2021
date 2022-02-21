package d5

import (
	"strings"
)

func P1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	segments := createSegments(lines)
	points := generatePoints(segments)
	intersections := findIntersections(points)
	return countIntersections(intersections)
}
