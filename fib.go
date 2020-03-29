package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(recursiveFib(i))
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(topdownFib(i))
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(bottomupFib(i))
	}
}

func recursiveFib(n int) int {
	if n < 2 {
		return n
	}

	return recursiveFib(n-1) + recursiveFib(n-2)
}

// topdownFib follows the top-down with memoization approach:
// - Solve the bigger problem by recursively solve the smaller subproblems
// - Whenever we solve a subproblem, we cache it and will then return the cache
//   if we see it again.
func topdownFib(n int) int {
	// memo stores calculated values, acts as a cache.
	memo := make([]int, n+1)
	return topdownFibHelper(memo, n)
}

func topdownFibHelper(memo []int, n int) int {
	if n < 2 {
		return n
	}

	// if we've already calculated it, simply return.
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = topdownFibHelper(memo, n-1) + topdownFibHelper(memo, n-2)
	return memo[n]
}

// bottomupFib follows the bottom-up with tabulation approach:
// - Fill up an n-dimensional table as we go up
func bottomupFib(n int) int {
	if n < 2 {
		return n
	}

	n1, n2, tmp := 0, 1, 0
	for i := 2; i < n+1; i++ {
		tmp = n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}
