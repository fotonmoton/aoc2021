package d15

import (
	"fmt"
	"strconv"
	"strings"
)

func expand(in [][]int) [][]int {
	height := len(in)
	width := len(in[0])

	expanded := make([][]int, height*5)

	for y := range expanded {
		expanded[y] = make([]int, width*5)
	}

	for y, line := range in {
		for x, num := range line {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					expanded[y+(height*i)][x+(width*j)] = (num + i + j) % 9
				}
			}
		}
	}

	return expanded
}

func expandMap(in string) string {
	lines := make([][]int, 0)
	for _, line := range strings.Split(strings.TrimSpace(in), "\n") {
		nums := make([]int, 0)
		for _, weight := range strings.Split(line, "") {
			weight, _ := strconv.Atoi(weight)
			nums = append(nums, weight)
		}
		lines = append(lines, nums)
	}

	expanded := expand(lines)

	newMap := ""

	for _, line := range expanded {
		for _, num := range line {
			newMap = newMap + fmt.Sprint(num)
		}

		newMap += "\n"
	}

	return newMap
}

func P2(in string) int {
	weights, neighbors := weightNeighbors(expandMap(in))
	return dijkstra(0, node(len(weights)-1), weights, neighbors)
}
