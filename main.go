package main

import (
	"fmt"
)

var XO_BOARD [3][3]string

func printBoard() {
	fmt.Println("-------")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Print("|")
			if XO_BOARD[x][y] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(XO_BOARD[x][y])
			}
		}
		fmt.Printf("|\n-------\n")
	}
	fmt.Println()
}

func scanInput() (row int, col int, inValid bool) {
	fmt.Print("row: ")
	fmt.Scanln(&row)
	row--

	fmt.Print("col: ")
	fmt.Scanln(&col)
	col--

	if row < 0 || row > 2 || col < 0 || col > 2 ||
		XO_BOARD[row][col] != "" {
		fmt.Println("Invalid input!")
		inValid = true
	}
	return
}

func isWinner() bool {
	if XO_BOARD[0][0] != "" && XO_BOARD[0][0] == XO_BOARD[1][1] && XO_BOARD[1][1] == XO_BOARD[2][2] ||
		XO_BOARD[0][2] != "" && XO_BOARD[0][2] == XO_BOARD[1][1] && XO_BOARD[1][1] == XO_BOARD[2][0] {
		return true
	}

	for i := 0; i < 3; i++ {
		if XO_BOARD[i][0] != "" && XO_BOARD[i][0] == XO_BOARD[i][1] && XO_BOARD[i][1] == XO_BOARD[i][2] ||
			XO_BOARD[0][i] != "" && XO_BOARD[0][i] == XO_BOARD[1][i] && XO_BOARD[1][i] == XO_BOARD[2][i] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("XO in Go!")
	printBoard()

	var currentPlayer string = "X"

	for numOfPlayes := 9; numOfPlayes > 0; {
		println("Player:", currentPlayer)
		row, col, inValid := scanInput()

		if inValid {
			continue
		}

		XO_BOARD[row][col] = currentPlayer

		printBoard()

		if isWinner() {
			fmt.Println(currentPlayer, "is the winner!")
			break
		}

		switch currentPlayer {
		case "X":
			currentPlayer = "O"
		case "O":
			currentPlayer = "X"
		}

		numOfPlayes--
	}

	fmt.Println("No one won!")
}
