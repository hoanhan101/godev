package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("lol"))
	fmt.Println(isPalindrome("loll"))
	fmt.Println(isPalindrome("llol"))
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
