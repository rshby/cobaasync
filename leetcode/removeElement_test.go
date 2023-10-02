package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	removeElement := func(nums []int, target int) int {
		i := 0
		n := len(nums)

		for i < n {
			if nums[i] == target {
				n--
				nums[i] = nums[n]
			} else {
				i++
			}
		}

		return n
	}

	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))

}
