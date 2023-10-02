package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSplitString(t *testing.T) {
	shortestPalindrome := func(s string) string {
		sliceOfWord := strings.Split(s, "")
		fmt.Println(sliceOfWord)
		fmt.Println("tipe data:", reflect.TypeOf(sliceOfWord))
		newSlice := []string{}

		for idx, _ := range sliceOfWord {
			newSlice = append(newSlice, sliceOfWord[len(sliceOfWord)-(idx+1)])
		}

		fmt.Println("reverse slice :", newSlice)

		polinderm := []string{}
		for _, item := range newSlice {
			polinderm = append(polinderm, item)
		}

		for _, item := range sliceOfWord {
			polinderm = append(polinderm, item)
		}

		newPolinderm := []string{}
		for i, value := range polinderm {
			if i != len(polinderm)-1 {
				if polinderm[i] == polinderm[i+1] {
					continue
				} else {
					newPolinderm = append(newPolinderm, value)
				}
			} else {
				newPolinderm = append(newPolinderm, value)
			}
		}

		fmt.Println("new polindersm :", strings.Join(newPolinderm, ""))

		return ""
	}

	shortestPalindrome("abcd")
}
