package d15

func P1(in string) int {
	weights, neighbors := weightNeighbors(in)
	return dijkstra(0, node(len(weights)-1), weights, neighbors)
}
