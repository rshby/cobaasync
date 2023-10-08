package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseInt(t *testing.T) {
	reverseInt := func(input int) int {
		flag := 1
		if input < 0 {
			flag = -1
			input *= -1
		}

		var temp int
		for input > 0 {
			temp = temp*10 + input%10
			input /= 10
		}

		return flag * temp
	}

	fmt.Println(reverseInt(-23))
}

func TestReverseInt2(t *testing.T) {
	reverseInt := func(input int) int {
		flag := 1
		if input < 0 {
			flag = -1
			input *= -1
		}

		temp := 0
		for input > 0 {
			temp = temp*10 + input%10
			input /= 10
		}

		return flag * temp
	}

	fmt.Println(reverseInt(234))
}
