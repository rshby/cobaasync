package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	searchInsert := func(input []int, target int) int {
		start := 0
		end := len(input) - 1
		var mid int

		for start <= end {
			mid = (start + end + 1) / 2
			if input[mid] == target {
				return mid
			}

			if input[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		return start
	}

	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, 2))
}
