package main

func setZeroes1(matrix [][]int) {
	col0 := 1
	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				matrix[i][0] = 0
				if j == 0 {
					col0 = 0
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				if matrix[i][0] == 0 || matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if col0 == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
