package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicat(t *testing.T) {
	isDuplicate := func(nums []int) bool {
		temp := map[int]int{}

		for id, value := range nums {
			_, ok := temp[value] // cek dulu jika ketemu value di temp
			if ok {
				return true
			}

			temp[value] = id // jika tidak ketemu, insert value ke temp
		}

		return false
	}

	fmt.Println(isDuplicate([]int{1, 1, 2, 3, 4}))
}
