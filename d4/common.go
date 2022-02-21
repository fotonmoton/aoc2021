package d4

import (
	"strconv"
	"strings"
)

type cell struct {
	value  int
	marked bool
}

type board []cell

const SIDE_SIZE = 5

func createBoards(rawBoards []string) []board {
	boards := make([]board, len(rawBoards))

	for i, rawBoard := range rawBoards {
		for _, cellRow := range strings.Split(rawBoard, "\n") {
			for _, rawCell := range strings.Split(cellRow, " ") {
				if rawCell != "" {
					value, _ := strconv.Atoi(rawCell)
					boards[i] = append(boards[i], cell{value, false})
				}
			}
		}
	}

	return boards
}

func markBoard(b board, take int) board {
	for i := range b {
		if b[i].value == take {
			b[i].marked = true
		}
	}

	return b
}

func gotBingo(b board) bool {

	for i := 0; i < SIDE_SIZE; i++ {
		if b[i*SIDE_SIZE].marked &&
			b[i*SIDE_SIZE+1].marked &&
			b[i*SIDE_SIZE+2].marked &&
			b[i*SIDE_SIZE+3].marked &&
			b[i*SIDE_SIZE+4].marked {
			return true
		}
	}

	for j := 0; j < SIDE_SIZE; j++ {
		if b[j].marked &&
			b[j+SIDE_SIZE].marked &&
			b[j+SIDE_SIZE*2].marked &&
			b[j+SIDE_SIZE*3].marked &&
			b[j+SIDE_SIZE*4].marked {
			return true
		}
	}

	return false
}

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

func score(b board, take int) int {
	sum := 0

	for _, cell := range b {
		if !cell.marked {
			sum += cell.value
		}
	}

	return sum * take
}
