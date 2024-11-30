package main

import "fmt"

func verticalCheck(board [3][3]int) int {
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[0][i] != 0 {
			return board[0][i]
		}

		if board[i][0] == board[i][1] && board[i][0] == board[i][2] && board[i][0] != 0 {
			return board[i][0]
		}
	}
	return 0
}

func diagonalCheck(board [3][3]int) int {
	if (board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != 0) ||
		(board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[0][2] != 0) {
		return board[0][0]
	}

	return 0
}

func IsSolved(board [3][3]int) int {
	if verticalCheck(board) != 0 {
		return verticalCheck(board)
	}
	if diagonalCheck(board) != 0 {
		return diagonalCheck(board)
	}

	for i := 0; i < 3; i++ {
		if board[i][0] == 0 || board[i][1] == 0 || board[i][2] == 0 {
			return -1
		}
	}

	return 0
}

func main() {
	board := [3][3]int{
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0},
	}
	board = [3][3]int{
		{2, 1, 2},
		{2, 2, 1},
		{1, 2, 2},
	}

	fmt.Println(IsSolved(board))
}