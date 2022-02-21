package d3

func rating(binaries []string, position int, path func([]string, []string) []string) string {
	if len(binaries) == 1 {
		return binaries[0]
	}

	ones := make([]string, 0, len(binaries)/2) // allocate at least half original slice
	zeroes := make([]string, 0, len(binaries)/2)
	for _, binary := range binaries {
		if binary[position] == '1' {
			ones = append(ones, binary)
		} else {
			zeroes = append(zeroes, binary)

		}
	}

	return rating(path(ones, zeroes), position+1, path)
}
