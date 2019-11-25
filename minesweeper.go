package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	m = 6   // rows
	n = 6   // columns
	p = 0.3 // probability of bombs
)

var (
	board    [m][n]string // map board
	solution [m][n]int    // solution board
)

func makeBoard() {
	rand.Seed(1)

	// put bombs (*) in the map.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if rand.Float64() < p {
				board[r][c] = "*"
			} else {
				board[r][c] = "."
			}
		}
	}
}

func solve() {
	// initialize solution board.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			solution[r][c] = 0
		}
	}

	// calculate solution.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			for rr := r - 1; rr <= r+1; rr++ {
				for cc := c - 1; cc <= c+2; cc++ {
					if (rr > 0 && rr < m) && (cc > 0 && cc < n) {
						if board[rr][cc] == "*" {
							solution[r][c]++
						}
					}
				}
			}
		}
	}

	// plug solution into board.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] != "*" {
				board[r][c] = strconv.Itoa(solution[r][c])
			}
		}
	}
}

func printBoard() {
	fmt.Println()
	for _, r := range board {
		fmt.Println(r)
	}
}

func printSolution() {
	fmt.Println()
	for _, r := range solution {
		fmt.Println(r)
	}
}

func main() {
	makeBoard()
	printBoard()

	solve()
	printBoard()
}
