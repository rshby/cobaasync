package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	isValid := func(s string) bool {
		stk := make([]byte, len(s))
		top := -1

		for i := range s {
			switch s[i] {
			case '(', '[', '{':
				stk[top+1] = s[i]
				top++
			case ')':
				if top == -1 || stk[top] != '(' {
					return false
				}
				top--
			case ']':
				if top == -1 || stk[top] != '[' {
					return false
				}
				top--
			case '}':
				if top == -1 || stk[top] != '{' {
					return false
				}
				top--
			}
		}

		return top == -1
	}

	fmt.Println(isValid("()[]"))
}
