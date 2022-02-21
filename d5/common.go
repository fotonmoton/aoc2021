package d5

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type segment struct {
	a point
	b point
}

func createSegments(lines []string) []segment {
	segments := make([]segment, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		segments[i] = segment{point{x1, y1}, point{x2, y2}}
	}

	return segments
}

func generatePoints(segments []segment) []point {
	points := []point{}

	for _, segment := range segments {
		if segment.a.x == segment.b.x {
			min := int(math.Min(float64(segment.a.y), float64(segment.b.y)))
			max := int(math.Max(float64(segment.a.y), float64(segment.b.y)))
			for i := min; i <= max; i++ {
				points = append(points, point{segment.a.x, i})
			}
		}

		if segment.a.y == segment.b.y {
			min := int(math.Min(float64(segment.a.x), float64(segment.b.x)))
			max := int(math.Max(float64(segment.a.x), float64(segment.b.x)))
			for i := min; i <= max; i++ {
				points = append(points, point{i, segment.a.y})
			}
		}
	}

	return points
}

func generatePointsWithDiagonals(segments []segment) []point {
	points := []point{}

	for _, segment := range segments {
		startX := segment.a.x
		startY := segment.a.y
		endX := segment.b.x
		endY := segment.b.y

		dx := 1
		dy := 1

		if segment.a.x == segment.b.x {
			dx = 0
		} else if segment.a.x > segment.b.x {
			dx = -1
		}

		if segment.a.y == segment.b.y {
			dy = 0
		} else if segment.a.y > segment.b.y {
			dy = -1
		}

		points = append(points, point{startX, startY})
		for startX != endX || startY != endY {
			startX += dx
			startY += dy
			points = append(points, point{startX, startY})
		}
	}

	return points
}

func findIntersections(points []point) map[point]int {

	intersections := map[point]int{}

	for _, a := range points {
		intersections[a]++
	}

	return intersections
}

func countIntersections(intersections map[point]int) int {
	count := 0

	for _, crosses := range intersections {
		if crosses > 1 {
			count++
		}
	}

	return count
}
