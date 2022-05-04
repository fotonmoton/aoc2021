package d12

func P1(input string) int {
	return countPaths(createMap(getPairs(getLines(input))), "start", make(map[string]int), true)
}
