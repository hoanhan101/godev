package main

import (
	"fmt"
)

func main() {
	fmt.Println(countCharacters("hello"))
}

func countCharacters(s string) map[string]int {
	m := map[string]int{}

	for _, c := range s {
		if _, ok := m[string(c)]; ok {
			m[string(c)]++
		} else {
			m[string(c)] = 1
		}
	}

	return m
}
