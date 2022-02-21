package d7

import "strconv"

func sumUp(to int) int {
	result := 0

	for from := 0; from <= to; from++ {
		result += from
	}

	return result
}

func toNum(in []string) []int {
	nums := make([]int, len(in))

	for idx, str := range in {
		num, _ := strconv.Atoi(str)
		nums[idx] = num
	}

	return nums
}
