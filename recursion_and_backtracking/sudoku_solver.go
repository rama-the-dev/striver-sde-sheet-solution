package main

func solveSudoku(board [][]byte) {
	// Algo
	// check for each place if it is empty
	// if it is empty then
	// check if we can put any no between 1-9
	// if yes
	// then call recursion again
	// if recursion returns true then return true
	// this would only happen when all elements are filled up
	// i.e. all places are filled with numbers 1-9
	// if no return false
	_ = recursion(board)
}

func canPutNumAtCurrentPosition(board [][]byte, row, col int, numToPut byte) bool {
	for i := 0; i < 9; i++ {
		// keeping row constant and column varying will check in all columns
		// if same no is present i.e. from left to right in same row
		if board[row][i] == numToPut {
			return false
		}
		if board[i][col] == numToPut {
			return false
		}
		// now we need to check if current board of size 3*3 also have same number
		// for this to find all indexes
		rowIndex := 3*(row/3) + (i / 3)
		colIndex := 3*(col/3) + (i % 3)
		if board[rowIndex][colIndex] == numToPut {
			return false
		}
	}
	return true
}

func recursion(board [][]byte) bool {
	for i, row := range board {
		for j, elem := range row {
			// check if current block is empty
			if elem == '.' {
				// check which no from 1-9 we can put in this place
				for k := 1; k <= 9; k++ {
					if canPutNumAtCurrentPosition(board, i, j, byte('0'+k)) {
						board[i][j] = byte('0' + k)
						// if recursion with current I returns true means
						// we got the solution and no need to find another solution
						// as only one solution is expected
						if solved := recursion(board); solved {
							return true
						}
						// since current I did not help in solving board
						// hence remove that no from board and try for next number
						board[i][j] = '.'
					}
				}
				// if we can not put any no in current empty shell
				// then we should backtrack
				return false
			}
			// in case current element is not empty
		}
	}
	return true
}

// func main() {
// 	sudoku := [][]byte{
// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
// 	solveSudoku(sudoku)
// 	fmt.Println(sudoku)
// }
