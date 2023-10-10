package leetcode

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	containsDuplicate := func(nums []int) bool {
		temp := map[int]int{}

		for idx, value := range nums {
			if _, ok := temp[value]; ok {
				return true
			}

			temp[value] = idx
		}

		return false
	}

	fmt.Println(containsDuplicate([]int{1, 2, 5, 3}))
}
