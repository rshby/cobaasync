package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestMedianArray(t *testing.T) {
	getMedian := func(a, b []int) float64 {
		x := a
		for _, value := range b {
			x = append(x, value)
		}

		sort.Ints(x)
		if len(x)%2 == 0 {
			return float64(x[len(x)/2-1]+x[len(x)/2]) / 2.0
		}
		return float64(x[len(x)/2])
	}
	fmt.Println(getMedian([]int{1, 2}, []int{3, 4}))
}
