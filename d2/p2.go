package d2

import (
	"log"
	"strconv"
	"strings"
)

func P2(input string) int {
	directions := strings.Split(strings.TrimSpace(input), "\n")
	x := 0
	y := 0
	aim := 0

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
			y += aim * value
			x += value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return x * y
}
