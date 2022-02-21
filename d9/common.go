package d9

import (
	"fmt"
	"strings"
)

const wall = '9'

func addWalls(lines []string) []string {
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%v%v%v", string(wall), line, string(wall))
	}
	nines := strings.Repeat(string(wall), len(lines[0]))
	return append([]string{nines}, append(lines, nines)...)
}
