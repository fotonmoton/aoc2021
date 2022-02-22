package d10

var pairs = map[rune]rune{'[': ']', '(': ')', '{': '}', '<': '>'}
var score1 = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var score2 = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
