package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fotonmoton/aoc2021/client"
	"github.com/fotonmoton/aoc2021/d1"
	"github.com/fotonmoton/aoc2021/d2"
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
}
