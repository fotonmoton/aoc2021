package d4

import (
	"strconv"
	"strings"
)

func findBingo2(boards []board, takes []string) (board, int) {
	bingos := make(map[*board]int)
	wins := []*board{}
	for _, take := range takes {
		converted, _ := strconv.Atoi(take)
		for boardIndex := range boards {
			board := &boards[boardIndex]
			if _, ok := bingos[board]; !ok {
				marked := markBoard(*board, converted)
				if gotBingo(marked) {
					bingos[board] = converted
					wins = append(wins, board)
				}
			}
		}
	}

	return *wins[len(wins)-1], bingos[wins[len(wins)-1]]
}

func P2(input string) int {
	lines := strings.Split(input, "\n\n")
	takes := strings.Split(lines[0], ",")
	boards := createBoards(lines[1:])
	bingo, take := findBingo2(boards, takes)

	return score(bingo, take)
}
