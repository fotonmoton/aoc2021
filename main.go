package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fotonmoton/aoc2021/client"
	"github.com/fotonmoton/aoc2021/d1"
	"github.com/fotonmoton/aoc2021/d2"
	"github.com/fotonmoton/aoc2021/d3"
	"github.com/fotonmoton/aoc2021/d4"
	"github.com/fotonmoton/aoc2021/d5"
	"github.com/fotonmoton/aoc2021/d6"
	"github.com/fotonmoton/aoc2021/d7"
	"github.com/fotonmoton/aoc2021/d8"
	"github.com/fotonmoton/aoc2021/d9"
)

func main() {

	session, err := os.ReadFile("session")

	if err != nil {
		log.Fatal(err)
	}

	client := client.NewClient(strings.TrimSpace(string(session)))

	fmt.Printf("day 1: %v\n", d1.D1(client.Day(1)))
	fmt.Printf("day 1 part 2: %v\n", d1.D1P2(client.Day(1)))
	fmt.Printf("day 2: %v\n", d2.D2(client.Day(2)))
	fmt.Printf("day 2 part 2: %v\n", d2.D2P2(client.Day(2)))
	fmt.Printf("day 3: %v\n", d3.D3(client.Day(3)))
	fmt.Printf("day 3 part 2: %v\n", d3.D3P2(client.Day(3)))
	fmt.Printf("day 4: %v\n", d4.D4(client.Day(4)))
	fmt.Printf("day 4 part 2: %v\n", d4.D4P2(client.Day(4)))
	fmt.Printf("day 5: %v\n", d5.D5(client.Day(5)))
	fmt.Printf("day 5 part 2: %v\n", d5.D5P2(client.Day(5)))
	fmt.Printf("day 6: %v\n", d6.D6(client.Day(6), 80))
	fmt.Printf("day 6 part 2: %v\n", d6.D6P2(client.Day(6), 256))
	fmt.Printf("day 7: %v\n", d7.D7(client.Day(7)))
	fmt.Printf("day 7 part 2: %v\n", d7.D7P2(client.Day(7)))
	fmt.Printf("day 8: %v\n", d8.D8(client.Day(8)))
	fmt.Printf("day 8 part 2: %v\n", d8.D8P2(client.Day(8)))
	fmt.Printf("day 9: %v\n", d9.D9(client.Day(9)))
	fmt.Printf("day 9 part 2: %v\n", d9.D9P2(client.Day(9)))
}
