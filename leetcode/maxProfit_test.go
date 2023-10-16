package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func TestBestTimeBuyAndSellStock(t *testing.T) {
	maxProfit := func(prices []int) int {
		min := math.MaxUint32
		res := 0

		for _, price := range prices {
			if price > min {
				if price-min > res {
					res = price - min
				}
			} else {
				min = price
			}
		}

		return res
	}

	fmt.Println(maxProfit([]int{100, 200, 300, 400, 500}))
}
