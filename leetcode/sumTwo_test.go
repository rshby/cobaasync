package leetcode

import (
	"fmt"
	"testing"
)

func TestSumTwo(t *testing.T) {
	sumTwo := func(nums []int, target int) []int {
		temp := map[int]int{}

		for idx, value := range nums {
			if res, ok := temp[target-value]; ok {
				return []int{res, idx}
			}

			temp[value] = idx
		}

		return []int{}
	}

	fmt.Println(sumTwo([]int{5, 2, 9, 10, 22}, 7))
}
