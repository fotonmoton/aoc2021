package d4

import "strings"

func P2(input string) int {
	lines := strings.Split(input, "\n\n")
	takes := strings.Split(lines[0], ",")
	boards := createBoards(lines[1:])
	bingo, take := findBingo2(boards, takes)

	return score(bingo, take)
}
