package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	removeDuplicate := func(nums []int) bool {
		temp := map[int]int{}

		for id, value := range nums {
			if _, ok := temp[value]; ok {
				return true
			}

			temp[value] = id
		}

		return false
	}

	fmt.Println(removeDuplicate([]int{1, 2, 3, 4, 1}))
}
