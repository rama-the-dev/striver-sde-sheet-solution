package main

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
*/
func rotate(matrix [][]int) {
	// matrix transpose across diagonal
	for i := range matrix {
		for j := i + 1; j < len(matrix[0]); j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
	// reverse rows
	for i, row := range matrix {
		l, r := 0, len(row)-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
