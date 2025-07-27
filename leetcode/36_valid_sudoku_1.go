package main

func isValidSudoku(board [][]byte) bool {
	for i, row := range board {
		for j, elem := range row {
			if elem != '.' && !isValidElement(board, i, j, elem) {
				return false
			}
		}
	}
	return true
}

func isValidElement(board [][]byte, row, col int, elem byte) bool {
	// check if same element is not present in
	// same row, same column and same board of 3*3
	for i := 0; i < 9; i++ {
		// if element present in row return false
		// while doing check skip current position element
		if i != col && board[row][i] == elem {
			return false
		}
		// if element present in same column return false
		if i != row && board[i][col] == elem {
			return false
		}
		// if element present in same 3*3 board then return false
		boxRow := 3*(row/3) + (i / 3)
		boxCol := 3*(col/3) + (i % 3)
		// one of the index are different than exact location of element
		if (boxRow != row || boxCol != col) && board[boxRow][boxCol] == elem {
			return false
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
// 	ans := isValidSudoku(sudoku)
// 	fmt.Println(ans)
// }
