package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseString([]string{"h", "e", "l", "l", "o"}))
}

func reverseString(s []string) []string {
	start := 0
	end := len(s) - 1

	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp

		start++
		end--
	}

	return s
}
