package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutRepeatChar(t *testing.T) {
	cekLongestSubstring := func(s string) int {
		start := -1
		lenMax := 0
		dict := [256]int{}

		for i, _ := range dict {
			dict[i] = -1
		}

		for id, value := range s {
			v := dict[value]
			if v > start {
				start = v
			}

			length := id - start
			if length > lenMax {
				lenMax = length
			}

			dict[value] = id
		}

		return lenMax
	}

	fmt.Println(cekLongestSubstring("abcabcbb"))
	fmt.Println(cekLongestSubstring("bbbbb"))
	fmt.Println(cekLongestSubstring("pwwkew"))
}
