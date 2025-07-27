package main

import (
	"strings"
)

func solveNQueens2(n int) [][]string {
	// Algo
	ans := [][]string{}

	// create empty board filled with dot
	// sine strings are immutable hence we can not have []string
	// as doing string[i] = Q will throw Compile error
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = []rune(strings.Repeat(".", n))
	}

	// in place of checking in all 3 direction for attacking queen
	// we can use hashmap to check if any direction have attacking queen
	left := make([]int, n)
	topLeft, bottomLeft := make([]int, 2*n-1), make([]int, 2*n-1)

	// run recursion with backtracking to find all the solutions
	recur2(&ans, board, left, topLeft, bottomLeft, 0, n)
	return ans
}

func copyBoardToAns2(board [][]rune, n int) []string {
	ans := make([]string, n)
	for i := 0; i < len(board); i++ {
		ans[i] = string(board[i])
	}
	return ans
}

func isSafeToPutQueen2(left, topLeft, bottomLeft []int, row, col, n int) bool {
	topLeftIndexToLook := n - 1 + col - row
	return left[row] != 1 && topLeft[topLeftIndexToLook] != 1 && bottomLeft[row+col] != 1
}

func recur2(ans *[][]string, board [][]rune, left, topLeft, bottomLeft []int, col, n int) {
	// exit condition
	if col == n {
		boardCopy := copyBoardToAns2(board, n)
		*ans = append(*ans, boardCopy)
		return
	}
	// for given column check for each row whether we can put the queen
	for row := 0; row < n; row++ {
		if isSafeToPutQueen2(left, topLeft, bottomLeft, row, col, n) {
			// put queen in current board position
			board[row][col] = 'Q'
			// update the hashes for all 3 direcitons
			left[row] = 1
			topLeft[n-1+col-row] = 1
			bottomLeft[row+col] = 1
			recur2(ans, board, left, topLeft, bottomLeft, col+1, n)
			// reverse the hashe updates for all 3 direcitons
			left[row] = 0
			topLeft[n-1+col-row] = 0
			bottomLeft[row+col] = 0
			// remove the queen from current board position
			// for backtracking
			board[row][col] = '.'
		}
	}
}

// func main() {
// 	ans := solveNQueens2(4)
// 	fmt.Println(ans)
// }
