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
}
