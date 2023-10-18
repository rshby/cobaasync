package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestPolindromWord(t *testing.T) {
	polindromWord := func(s string) bool {
		word := s

		temp := []string{}
		for i := 0; i < len(s); i++ {
			temp = append(temp, string(s[len(s)-1-i]))
		}

		return word == strings.Join(temp, "")
	}

	fmt.Println(polindromWord("aba"))
}
