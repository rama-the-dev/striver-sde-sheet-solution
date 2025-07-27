package main

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	// Algo
	ans := [][]string{}

	// create empty board filled with dot
	// sine strings are immutable hence we can not have []string
	// as doing string[i] = Q will throw Compile error
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = []rune(strings.Repeat(".", n))
	}

	// run recursion with backtracking to find all the solutions
	recur(&ans, board, 0, n)
	return ans
}

func copyBoardToAns(board [][]rune, n int) []string {
	ans := make([]string, n)
	for i := 0; i < len(board); i++ {
		ans[i] = string(board[i])
	}
	return ans
}

func isSafeToPutQueen(board [][]rune, row, col, n int) bool {
	// out of total 8 directions only in 3 directions we have to check
	// 1. top left 2. left 3. bottom left
	rowCopy, colCopy := row, col

	// check in top left
	// both row and column would reduce and go upto zero
	for row >= 0 && col >= 0 {
		if board[row][col] == 'Q' {
			return false
		}
		row--
		col--
	}

	// check in left
	// only col value would reduce and row will be same
	row, col = rowCopy, colCopy
	for col >= 0 {
		if board[row][col] == 'Q' {
			return false
		}
		col--
	}

	// check in bottom left
	// row value would increase and col value would decrease
	row, col = rowCopy, colCopy
	for row < n && col >= 0 {
		if board[row][col] == 'Q' {
			return false
		}
		row++
		col--
	}
	return true
}

func recur(ans *[][]string, board [][]rune, col, n int) {
	// exit condition
	if col == n {
		boardCopy := copyBoardToAns(board, n)
		*ans = append(*ans, boardCopy)
		return
	}
	// for given column check for each row whether we can put the queen
	for row := 0; row < n; row++ {
		if isSafeToPutQueen(board, row, col, n) {
			board[row][col] = 'Q'
			recur(ans, board, col+1, n)
			board[row][col] = '.'
		}
	}
}

// func main() {
// 	ans := solveNQueens(4)
// 	fmt.Println(ans)
// }
