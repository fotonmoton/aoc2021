package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fotonmoton/aoc2021/client"
	"github.com/fotonmoton/aoc2021/d1"
	"github.com/fotonmoton/aoc2021/d2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your session: ")
	session, _ := reader.ReadString('\n')
	client := client.NewClient(strings.TrimSpace(session))

	fmt.Printf("day1: %v\n", d1.D1(string(client.Day(1))))
	fmt.Printf("day2: %v\n", d2.D2(string(client.Day(2))))
}
