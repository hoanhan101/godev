/*
simple 3x3 tic tac toe for 2.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var board [3][3]string
var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	initBoard()
	// place(0, 0, "X")
	// place(1, 1, "X")
	// place(2, 2, "X")

	start()
}

func check3(s1, s2, s3 string) bool {
	if s1 != "_" && s2 != "_" && s3 != "_" && s1 == s2 && s2 == s3 {
		return true
	}

	return false
}

func checkWin() bool {
	// check 9 win states.
	if check3(board[0][0], board[0][1], board[0][2]) ||
		check3(board[1][0], board[1][1], board[1][2]) ||
		check3(board[2][0], board[2][1], board[2][2]) ||
		check3(board[0][0], board[1][0], board[2][0]) ||
		check3(board[0][1], board[1][1], board[2][1]) ||
		check3(board[0][2], board[1][2], board[2][2]) ||
		check3(board[0][0], board[1][1], board[2][2]) ||
		check3(board[2][0], board[1][1], board[0][2]) {
		fmt.Println("yo win")
		return true
	}

	return false
}

func start() {
	for i := 0; i < 9; i++ {
		readMove()
	}
}

func readMove() {
	if scanner.Scan() {
		in := scanner.Text()

		// skip err checking here.
		r, _ := strconv.Atoi(string(in[0]))
		c, _ := strconv.Atoi(string(in[1]))
		s := string(in[2])
		place(r, c, s)
	}
}

func initBoard() {
	for r := range board {
		for c := range board {
			board[r][c] = "_"
		}
	}

	printBoard()
}

func printBoard() {
	for r := range board {
		for c := range board {
			fmt.Printf("%v ", board[r][c])
		}
		fmt.Println()
	}

	fmt.Println()
}

func place(r, c int, s string) {
	board[r][c] = s

	printBoard()

	checkWin()
}