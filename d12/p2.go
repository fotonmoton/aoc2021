package d12

func P2(input string) int {
	return countPaths(createMap(getPairs(getLines(input))), "start", make(map[string]int), false)
}
