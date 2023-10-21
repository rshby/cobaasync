package leetcode

import (
	"log"
	"math"
	"testing"
)

func TestBestTimeToSellAndBuy(t *testing.T) {
	maxProfit := func(prices []int) int {
		min := math.MaxUint32
		res := 0

		for _, price := range prices {
			if price > min {
				if price-min > res {
					res = price - min
				}
			} else { // price < min
				min = price
			}
		}

		return res
	}

	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
