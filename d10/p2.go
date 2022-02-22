package d10

import (
	"sort"
	"strings"
)

func findScore2(line string) int {
	total := 0
	stack := []rune{}
	for _, current := range line {
		if _, isLeft := pairs[current]; isLeft {
			stack = append(stack, current)
			continue
		}

		lastLeft := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current != pairs[lastLeft] {
			return -1
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		total = total*5 + score2[pairs[stack[i]]]
	}

	return total
}

func P2(in string) int {
	scores := []int{}
	for _, line := range strings.Split(in, "\n") {
		if score := findScore2(line); score != -1 {
			scores = append(scores, score)
		}
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	return scores[len(scores)/2]
}
