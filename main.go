package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	currentPlayer := "X"
	winner := ""

	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("To make a move, enter the row and column numbers separated by a space.")

	for winner == "" {
		printBoard(board)
		fmt.Printf("Player %s's turn: ", currentPlayer)

		var row, col int
		fmt.Scanf("%d %d", &row, &col)

		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid input. Please enter row and column numbers between 1 and 3.")
			continue
		}

		if board[row-1][col-1] != " " {
			fmt.Println("That spot is already taken. Please choose another.")
			continue
		}

		board[row-1][col-1] = currentPlayer

		if checkWinner(board, currentPlayer) {
			winner = currentPlayer
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}

	printBoard(board)
	fmt.Printf("Player %s wins!\n", winner)
}

func printBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Println(strings.Join(board[i][:], " | "))
		if i < 2 {
			fmt.Println("---------")
		}
	}
}

func checkWinner(board [3][3]string, player string) bool {
	// check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}
