package main

/*
Important points
1. map 2d matrix into 1d using below concept
	a. 1d matrix matrix last index = low * mid - 1
	b. 1d matrix mid elem row = mid / no of columns
	c. 1d matrix mid elem col = mid % no of columns
*/

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1

	for low <= high {
		mid := (low + high) / 2
		row, col := mid/n, mid%n
		if matrix[row][col] == target {
			return true
		} else if target < matrix[row][col] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

// func main() {
// 	matrix := [][]int{{1, 3}}
// 	fmt.Println(searchMatrix2(matrix, 3))
// }
