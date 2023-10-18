package leetcode

import (
	"fmt"
	"testing"
)

func TestPolindromNumber(t *testing.T) {
	polindromNumber := func(num int) bool {
		number := num

		var temp int
		for num > 0 {
			temp = temp*10 + num%10
			num /= 10
		}

		return number == temp
	}

	fmt.Println(polindromNumber(121))
}
