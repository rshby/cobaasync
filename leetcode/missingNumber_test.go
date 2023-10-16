package leetcode

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	missingNumber := func(nums []int) int {
		total := 0
		for _, value := range nums {
			total += value
		}

		l := len(nums)
		sum := l * (l + 1) / 2
		return sum - total
	}

	fmt.Println(missingNumber([]int{1, 2, 3, 5}))
}
