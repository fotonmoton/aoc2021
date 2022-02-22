package d10

import "strings"

func findScore1(line string) int {
	stack := []rune{}
	for _, current := range line {
		if _, isLeft := pairs[current]; isLeft {
			stack = append(stack, current)
			continue
		}

		lastLeft := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current != pairs[lastLeft] {
			return score1[current]
		}

	}
	return 0
}

func P1(in string) int {
	score := 0
	for _, line := range strings.Split(in, "\n") {
		score += findScore1(line)
	}
	return score
}
