package d4

import (
	"strconv"
	"strings"
)

func findBingo(boards []board, takes []string) (board, int) {

	for _, take := range takes {
		converted, _ := strconv.Atoi(take)
		for _, board := range boards {
			marked := markBoard(board, converted)
			if gotBingo(marked) {
				return marked, converted
			}
		}
	}

	return nil, 0
}

func P1(input string) int {
	lines := strings.Split(input, "\n\n")
	takes := strings.Split(lines[0], ",")
	boards := createBoards(lines[1:])
	bingo, take := findBingo(boards, takes)

	return score(bingo, take)
}
