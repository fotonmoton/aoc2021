package d2

import (
	"log"
	"strconv"
	"strings"
)

type position struct {
	x int64
	y int64
}

func P1(input string) int64 {
	position := position{}
	directions := strings.Split(strings.TrimSpace(input), "\n")

	for _, direction := range directions {
		directionAndValue := strings.Split(strings.TrimSpace(direction), " ")

		if len(directionAndValue) != 2 {
			log.Fatal("line has wrong number of elements: ", direction)
		}

		value, err := strconv.Atoi(directionAndValue[1])

		if err != nil {
			log.Fatal(err)
		}

		switch directionAndValue[0] {
		case "forward":
			position.x += int64(value)
		case "down":
			position.y += int64(value)
		case "up":
			position.y -= int64(value)
		}
	}

	return position.x * position.y
}
