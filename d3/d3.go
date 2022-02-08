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
