package leetcode

import (
	"log"
	"sort"
	"testing"
)

func TestMedianArray(t *testing.T) {
	medianArray := func(a, b []int) float64 {
		temp := a
		for _, value := range b {
			temp = append(temp, value)
		}

		sort.Ints(temp)
		if len(temp)%2 == 0 {
			return float64(temp[len(temp)/2-1]+temp[len(temp)/2]) / 2.0
		}

		return float64(temp[len(temp)/2])
	}

	log.Println(medianArray([]int{1, 2, 3, 4}, []int{5, 6}))
}
