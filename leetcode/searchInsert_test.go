package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	searchInsert := func(nums []int, target int) int {
		start := 0
		end := len(nums) - 1
		var mid int

		for start <= end {
			mid = (start + end + 1) / 2
			if nums[mid] == target {
				return mid
			}

			if nums[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		return start
	}

	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5, 6, 7, 8}, 0))
}
