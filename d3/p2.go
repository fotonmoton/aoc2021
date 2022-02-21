package d3

import (
	"strconv"
	"strings"
)

func P2(input string) int64 {
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
