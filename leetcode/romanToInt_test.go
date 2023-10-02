package leetcode

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	romanToInt := func(s string) int {
		m := map[byte]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		result := m[s[0]]
		for i := 1; i < len(s); i++ {
			last, now := m[s[i-1]], m[s[i]]
			if last < now {
				result += now - 2*last
			} else {
				result += now
			}
		}
		return result
	}

	fmt.Println(romanToInt("XX"))
}
