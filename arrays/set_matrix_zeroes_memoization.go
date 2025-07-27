package main

/*
73. Set Matrix Zeroes
Solved
Medium
Topics
Companies
Hint
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	rowZeroMem := make([]int, rows)
	colZeroMem := make([]int, cols)

	// memorize 0 for row and column index
	for i, row := range matrix {
		for j, elem := range row {
			if (elem) == 0 {
				rowZeroMem[i] = 1
				colZeroMem[j] = 1
			}
		}
	}

	// set matrix 0's based on row and column zero memoization
	for i, row := range matrix {
		for j := range row {
			if rowZeroMem[i] == 1 || colZeroMem[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
