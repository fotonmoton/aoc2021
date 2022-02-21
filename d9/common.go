package d9

import (
	"fmt"
	"strings"
)

func addWalls(lines []string) []string {
	for i, line := range lines {
		lines[i] = fmt.Sprintf("9%v9", line)
	}

	nines := strings.Repeat("9", len(lines[0]))
	return append([]string{nines}, append(lines, nines)...)
}
