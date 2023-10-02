package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutRepeat(t *testing.T) {
	cekLongestSubstring := func(s string) int {
		start := -1
		lenMax := 0
		dict := [256]int{}

		for i, _ := range dict {
			dict[i] = -1
		}

		for idx, value := range s {
			v := dict[value]
			if v > start {
				start = v
			}

			length := idx - start
			if length > lenMax {
				lenMax = length
			}

			dict[value] = idx
		}

		return lenMax
	}

	fmt.Println(cekLongestSubstring("abcabcbb"))
	fmt.Println(cekLongestSubstring("bbbbb"))
	fmt.Println(cekLongestSubstring("pwwkew"))
}
