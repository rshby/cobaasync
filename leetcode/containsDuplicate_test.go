package leetcode

import (
	"fmt"
	"testing"
)

func TestContainDuplicate(t *testing.T) {
	containDup := func(nums []int) bool {
		temp := map[int]int{}

		for id, value := range nums {
			if _, ok := temp[value]; ok {
				return true
			}

			temp[value] = id
		}

		return false
	}

	fmt.Println(containDup([]int{1, 2, 3, 4, 5}))
}
