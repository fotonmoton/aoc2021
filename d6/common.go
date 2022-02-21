package d6

import "strconv"

func toNum(in []string) []int {
	nums := make([]int, len(in))

	for idx, str := range in {
		num, _ := strconv.Atoi(str)
		nums[idx] = num
	}

	return nums
}
