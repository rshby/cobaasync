package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortMedian(t *testing.T) {
	findMedian := func(a, b []int) float64 {
		temp := a
		for _, value := range b {
			temp = append(temp, value)
		}

		sort.Ints(temp)

		if len(temp)%2 == 0 {
			return float64(temp[(len(temp)/2)-1]+temp[len(temp)/2]) / 2.0
		}

		return float64(temp[len(temp)/2])
	}

	fmt.Println(findMedian([]int{1, 2, 3}, []int{4, 5}))
}
