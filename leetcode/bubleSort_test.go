package leetcode

import (
	"log"
	"testing"
)

func TestBubleSorting(t *testing.T) {
	bubleSort := func(nums []int) []int {
		isDone := false

		for !isDone {
			isDone = true
			for i := 0; i < len(nums)-1; i++ {
				if nums[i] > nums[i+1] {
					nums[i], nums[i+1] = nums[i+1], nums[i]
					isDone = false
				}
			}
		}

		return nums
	}

	log.Println(bubleSort([]int{5, 1, 2, 3, 4}))
}
