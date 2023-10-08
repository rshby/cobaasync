package leetcode

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	isAnagram := func(s, t string) bool {
		if len(s) != len(t) {
			return false
		}

		var temp = [26]int{}

		for idx := 0; idx < len(s); idx++ {
			temp[s[idx]-'a']++
			temp[t[idx]-'a']--
		}

		for i := 0; i < len(temp); i++ {
			if temp[i] != 0 {
				return false
			}
		}

		return true
	}

	fmt.Println(isAnagram("anagram", "nagaram"))
}
