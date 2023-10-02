package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthLongestSubstr(t *testing.T) {
	lengthOfLongestSubstring := func(s string) int {
		start := -1
		lenMax := 0
		dict := [256]int{}
		for i := range dict {
			dict[i] = -1
		}
		//fmt.Println("dict :", dict)
		for i, value := range s { // ada 7 perulangan
			if v := dict[value]; v > start { // -1 > -1 -> false di perulangan pertama
				start = v
			}
			length := i - start // 0 - -1 = 1
			if length > lenMax {
				lenMax = length
			}
			dict[value] = i
		}
		return lenMax
	}

	x := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(x)
}
