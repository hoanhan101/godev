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

	c := make(chan bool, 1)
	go func() {
		play(c)
	}()

	done := <-c
	fmt.Println("complete?", done)
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

func play(complete chan bool) {
	for {
		done := move()
		if done {
			complete <- true
			close(complete)
			return
		}
	}
}

func move() bool {
	if scanner.Scan() {
		in := scanner.Text()

		// skip err checking here.
		r, _ := strconv.Atoi(string(in[0]))
		c, _ := strconv.Atoi(string(in[1]))
		s := string(in[2])
		place(r, c, s)

		return checkWin()
	}

	return false
}

func place(r, c int, s string) {
	board[r][c] = s
	printBoard()
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
		return true
	}

	return false
}

func check3(s1, s2, s3 string) bool {
	if s1 != "_" && s2 != "_" && s3 != "_" && s1 == s2 && s2 == s3 {
		return true
	}

	return false
}
