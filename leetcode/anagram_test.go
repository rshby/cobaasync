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

		temp := [26]int{}

		for i := 0; i < len(s); i++ {
			temp[s[i]-'a']++
			temp[t[i]-'a']--
		}

		for _, value := range temp {
			if value != 0 {
				return false
			}
		}

		return true
	}

	fmt.Println(isAnagram("anagram", "nagaram"))
}
