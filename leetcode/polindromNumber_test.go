package leetcode

import (
	"fmt"
	"testing"
)

func TestPolindromNumber(t *testing.T) {
	cekPolindrom := func(x int) bool {
		number := x

		temp := 0
		for x > 0 {
			temp = temp*10 + x%10
			x /= 10
		}

		return number == temp
	}

	fmt.Println(cekPolindrom(122))
}
