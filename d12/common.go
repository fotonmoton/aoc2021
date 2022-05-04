package d12

import (
	"strings"
)

type caveMap map[string][]string

func getLines(in string) []string {
	return strings.Split(strings.TrimSpace(in), "\n")
}

func getPairs(lines []string) [][]string {
	pairs := [][]string{}

	for _, line := range lines {
		pairs = append(pairs, strings.Split(line, "-"))
	}

	return pairs
}

func createMap(paths [][]string) caveMap {
	caveMap := make(map[string][]string)

	for _, path := range paths {
		left := path[0]
		right := path[1]

		caveMap[left] = append(caveMap[left], right)
		caveMap[right] = append(caveMap[right], left)
	}

	return caveMap
}

func countPaths(caveMap caveMap, key string, visited map[string]int, duplicate bool) int {
	count := 0

	if key == "end" {
		return 1
	}

	if visited[key] > 0 {
		if key == "start" {
			return 0
		}

		if strings.ToUpper(key) != key {
			if duplicate {
				return 0
			} else {
				duplicate = true
			}
		}
	}

	visited[key]++
	for _, path := range caveMap[key] {
		count += countPaths(caveMap, path, visited, duplicate)
	}
	visited[key]--

	return count
}
