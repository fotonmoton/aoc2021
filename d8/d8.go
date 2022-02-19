package d8

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func D8(in string) int {
	lines := strings.Split(strings.TrimSpace(in), "\n")

	sum := 0

	for _, line := range lines {
		for _, digit := range strings.Split(strings.Split(line, "|")[1], " ") {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}
	}

	return sum
}

func difference(a []string, b []string) []string {
	needle := a
	haystack := b

	if len(a) > len(b) {
		needle = b
		haystack = a
	}

	mb := make(map[string]struct{}, len(needle))
	for _, x := range needle {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range haystack {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func merge(a []string, b []string) []string {
	return append(append([]string{}, a...), b...)
}

func findSegment(digits [][]string, part []string) string {
	for _, digit := range digits {
		diff := difference(digit, part)

		if len(digit) > len(part) && len(diff) == 1 {
			return diff[0]
		}
	}

	log.Fatal("can't find segment")

	return ""
}

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg
//
// a: 0, c: 1, f: 2, g: 3, e: 4, b: 5, d: 6
func findDisplay(digitsAsString string) []string {
	display := make([]string, 7)

	digits := [][]string{}
	splitted := strings.Split(digitsAsString, " ")
	sort.Slice(splitted, func(i, j int) bool { return len(splitted[i]) < len(splitted[j]) })
	for _, digit := range splitted {
		digits = append(digits, strings.Split(digit, ""))
	}

	// a
	display[0] = findSegment(digits, digits[0])
	// g
	display[3] = findSegment(digits, merge(digits[2], []string{display[0]}))
	// e
	display[4] = findSegment(digits, merge(digits[2], []string{display[0], display[3]}))
	// d
	display[6] = findSegment(digits, merge(digits[1], []string{display[3]}))
	// b
	display[5] = findSegment(digits, merge(digits[1], []string{display[3], display[4], display[6]}))
	// f
	display[2] = findSegment(digits, []string{display[0], display[3], display[4], display[5], display[6]})
	// c
	display[1] = findSegment(digits, []string{display[0], display[2], display[3], display[4], display[5], display[6]})

	return display
}

func isOne(digit []string, display []string) bool {
	return len(digit) == 2
}

func isTwo(digit []string, display []string) bool {

	if len(digit) != 5 {
		return false
	}

	two := strings.Join([]string{display[0], display[1], display[3], display[4], display[6]}, "")

	for _, segment := range digit {
		if !strings.Contains(two, segment) {
			return false
		}
	}

	return true
}

func isThree(digit []string, display []string) bool {

	if len(digit) != 5 {
		return false
	}

	three := strings.Join([]string{display[0], display[1], display[2], display[3], display[6]}, "")

	for _, segment := range digit {
		if !strings.Contains(three, segment) {
			return false
		}
	}

	return true
}

func isFour(digit []string, display []string) bool {
	return len(digit) == 4
}

func isFive(digit []string, display []string) bool {

	if len(digit) != 5 {
		return false
	}

	three := strings.Join([]string{display[0], display[2], display[3], display[5], display[6]}, "")

	for _, segment := range digit {
		if !strings.Contains(three, segment) {
			return false
		}
	}

	return true
}

func isSix(digit []string, display []string) bool {

	if len(digit) != 6 {
		return false
	}

	three := strings.Join([]string{display[0], display[2], display[3], display[4], display[5], display[6]}, "")

	for _, segment := range digit {
		if !strings.Contains(three, segment) {
			return false
		}
	}

	return true
}

func isSeven(digit []string, display []string) bool {
	return len(digit) == 3
}

func isEight(digit []string, display []string) bool {
	return len(digit) == 7
}

func isNine(digit []string, display []string) bool {

	if len(digit) != 6 {
		return false
	}

	three := strings.Join([]string{display[0], display[1], display[2], display[3], display[5], display[6]}, "")

	for _, segment := range digit {
		if !strings.Contains(three, segment) {
			return false
		}
	}

	return true
}

func isZero(digit []string, display []string) bool {

	if len(digit) != 6 {
		return false
	}

	three := strings.Join([]string{display[0], display[1], display[2], display[3], display[4], display[5]}, "")

	for _, segment := range digit {
		if !strings.Contains(three, segment) {
			return false
		}
	}

	return true
}

func digitToNumber(digit []string, display []string) int {
	numberMap := map[int]func([]string, []string) bool{
		0: isZero,
		1: isOne,
		2: isTwo,
		3: isThree,
		4: isFour,
		5: isFive,
		6: isSix,
		7: isSeven,
		8: isEight,
		9: isNine,
	}

	for number, f := range numberMap {
		if f(digit, display) {
			return number
		}
	}
	log.Fatal("wrong digit", digit, display)
	return 0
}

func findCode(digitsAsString string, display []string) int {
	numbers := []string{}
	for _, digit := range strings.Split(digitsAsString, " ") {
		numbers = append(numbers, strconv.Itoa(digitToNumber(strings.Split(digit, ""), display)))
	}
	asInt, _ := strconv.Atoi(strings.Join(numbers, ""))

	return asInt
}

func D8P2(in string) int {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	sum := 0
	for _, line := range lines {
		wiringsAndCode := strings.Split(line, " | ")
		display := findDisplay(wiringsAndCode[0])
		sum += findCode(wiringsAndCode[1], display)
	}

	return sum
}
