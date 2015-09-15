package todoutil

import (
	"strconv"
)

func Atois(words []string) ([]int, error) {
	var nums []int
	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func ContainsInt(xs []int, n int) bool {
	for _, x := range xs {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsString(strs []string, str string) bool {
	for _, x := range strs {
		if x == str {
			return true
		}
	}
	return false
}
