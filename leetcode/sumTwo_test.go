package leetcode

import (
	"fmt"
	"testing"
)

func TestSumTwo(t *testing.T) {
	sumTwo := func(nums []int, target int) []int {
		temp := map[int]int{}

		for idx, value := range nums {
			hasilKurang := target - value
			res, ok := temp[hasilKurang]
			if ok {
				return []int{res, idx}
			}

			temp[value] = idx
		}

		return []int{}
	}

	fmt.Println(sumTwo([]int{2, 5, 7, 8, 9}, 9))
}
