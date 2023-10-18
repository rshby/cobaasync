package leetcode

import (
	"log"
	"testing"
)

func TestSumTwo(t *testing.T) {
	sumTwo := func(nums []int, target int) []int {
		temp := map[int]int{}

		for id, value := range nums {
			if res, ok := temp[target-value]; ok {
				return []int{res, id}
			}

			temp[value] = id
		}

		return []int{}
	}

	log.Println(sumTwo([]int{1, 3, 5, 2, 7}, 9))
}
