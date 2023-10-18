package leetcode

import (
	"fmt"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	maxArea := func(height []int) int {
		var max, left, t int
		right := len(height) - 1

		for left < right {
			if height[left] < height[right] {
				t = height[left] * (right - left)
				left++
			} else {
				t = height[right] * (right - left)
				right--
			}

			if t > max {
				max = t
			}
		}

		return max
	}

	fmt.Println(maxArea([]int{100, 200, 300, 400}))
}
