package d5

import "strings"

func P2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	segments := createSegments(lines)
	points := generatePointsWithDiagonals(segments)
	intersections := findIntersections(points)
	return countIntersections(intersections)
}
