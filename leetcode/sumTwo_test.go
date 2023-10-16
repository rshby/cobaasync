package leetcode

import (
	"fmt"
	"testing"
)

func TestSumTwo(t *testing.T) {
	sumTwo := func(nums []int, target int) []int {
		temp := map[int]int{}

		for id, value := range nums {
			if result, ok := temp[target-value]; ok {
				return []int{result, id}
			}

			temp[value] = id
		}

		return []int{}
	}

	fmt.Println(sumTwo([]int{5, 2, 7, 10}, 9))
}
