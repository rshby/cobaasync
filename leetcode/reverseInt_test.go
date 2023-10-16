package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseInt(t *testing.T) {
	reverseInt := func(num int) int {
		flag := 1
		if num < 0 {
			flag = -1
			num *= -1
		}

		temp := 0

		for num > 0 {
			temp = temp*10 + num%10
			num /= 10
		}

		return flag * temp
	}

	fmt.Println(reverseInt(123))
}
