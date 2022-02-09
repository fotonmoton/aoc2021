package d3

import (
	"strconv"
	"strings"
)

func D3(input string) int64 {
	binaries := strings.Split(strings.TrimSpace(input), "\n")

	ones := make([]int, len(binaries[0]))

	for _, binary := range binaries {
		for position, bit := range binary {
			if byte(bit) == '1' {
				ones[position]++
			}
		}
	}

	gamma := make([]string, len(ones))
	epsilon := make([]string, len(ones))

	for pos, oneCount := range ones {
		if oneCount > len(binaries)/2 {
			gamma[pos] = "1"
			epsilon[pos] = "0"
		} else {
			gamma[pos] = "0"
			epsilon[pos] = "1"
		}
	}

	gammaNum, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilonNum, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	return gammaNum * epsilonNum
}

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

func D3P2(input string) int64 {
	binaries := strings.Split(strings.TrimSpace(input), "\n")

	o2predicate := func(ones, zeroes []string) []string {
		if len(ones) == len(zeroes) || len(ones) > len(zeroes) {
			return ones
		}
		return zeroes
	}

	co2predicate := func(ones, zeroes []string) []string {
		if len(ones) == len(zeroes) || len(ones) > len(zeroes) {
			return zeroes
		}
		return ones
	}

	o2rating, _ := strconv.ParseInt(rating(binaries, 0, o2predicate), 2, 64)
	co2rating, _ := strconv.ParseInt(rating(binaries, 0, co2predicate), 2, 64)

	return o2rating * co2rating
}
