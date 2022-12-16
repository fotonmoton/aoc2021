package d15

import (
	"math"

	"golang.org/x/exp/maps"
)

type node int

type weights map[node]int

type neighbors map[node][]node

func minDistance(unvisited weights, distances weights) node {
	var min node
	var minDistance = math.MaxInt64
	for current := range unvisited {
		if distances[current] < minDistance {
			min = current
			minDistance = distances[current]
		}
	}

	return min
}

func assignMap[M ~map[K]V, K comparable, V any](m M, v V) M {
	n := make(M, len(m))
	for k := range m {
		n[k] = v
	}
	return n

}

func dijkstra(start node, end node, weights weights, neighbors neighbors) int {
	distances := assignMap(weights, math.MaxInt64)
	unvisited := maps.Clone(weights)

	distances[start] = 0

	for len(unvisited) != 0 {
		current := minDistance(unvisited, distances)

		delete(unvisited, current)

		for _, neighbor := range neighbors[current] {
			oldDistance := distances[neighbor]
			newDistance := weights[neighbor] + distances[current]
			if newDistance < oldDistance {
				distances[neighbor] = newDistance
			}
		}
	}

	return distances[end]
}
