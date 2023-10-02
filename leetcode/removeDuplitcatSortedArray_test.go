package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateArray(t *testing.T) {
	removeDuplicate := func(nums []int) int {
		length := len(nums)
		newLength := 1

		for i := 1; i < length; i++ {
			if nums[i-1] != nums[i] {
				nums[newLength] = nums[i]
				newLength++
			}
		}

		return newLength
	}

	fmt.Println(removeDuplicate([]int{1, 1, 2}))
}
