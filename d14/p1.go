package d14

import (
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func reduce[S any, A any](slice []S, f func(elm S, acc A) A, acc A) A {
	for _, elm := range slice {
		acc = f(elm, acc)
	}
	return acc
}

func makePairs(template string) map[string]int {
	pairs := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])] = 1
	}

	return pairs
}

func templateAndInstructions(in string) (string, map[string]string) {
	parts := strings.Split(strings.TrimSpace(in), "\n\n")
	instructions := make(map[string]string)
	for _, pair := range strings.Split(parts[1], "\n") {
		mapping := strings.Split(pair, " -> ")
		instructions[mapping[0]] = mapping[1]
	}
	return parts[0], instructions
}

func count(template string, instructions map[string]string, runs int) map[string]int {
	init := func(char string, acc map[string]int) map[string]int { acc[char]++; return acc }
	counts := reduce(strings.Split(template, ""), init, make(map[string]int))
	pairs := makePairs(template)
	for i := 0; i < runs; i++ {
		for pair, count := range maps.Clone(pairs) {
			left := pair[:1]
			right := pair[1:]
			middle := instructions[pair]
			pairs[pair] -= count
			pairs[left+middle] += count
			pairs[middle+right] += count
			counts[middle] += count
		}
	}

	return counts
}

func countPolymer(in string, runs int) int {
	template, instructions := templateAndInstructions(in)
	counts := maps.Values(count(template, instructions, runs))
	slices.Sort(counts)
	max := counts[len(counts)-1]
	min := counts[0]

	return max - min
}

func P1(in string) int {
	return countPolymer(in, 10)
}
