package leetcode

import (
	"fmt"
	"testing"
)

func TestBubleSort(t *testing.T) {
	bubleSort := func(nums []int) []int {
		isDone := false

		for !isDone {
			isDone = true
			i := 0
			for i < len(nums)-1 {
				if nums[i] > nums[i+1] {
					nums[i], nums[i+1] = nums[i+1], nums[i]
					isDone = false
				}
				i++
			}
		}
		return nums
	}
	fmt.Println(bubleSort([]int{4, 6, 5, 1, 2}))
}
