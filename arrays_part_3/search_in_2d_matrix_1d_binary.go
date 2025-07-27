package main

/*
	You are given an m x n integer matrix matrix with the following two properties:

	Each row is sorted in non-decreasing order.
	The first integer of each row is greater than the last integer of the previous row.
	Given an integer target, return true if target is in matrix or false otherwise.

	You must write a solution in O(log(m * n)) time complexity.
*/

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for _, row := range matrix {
		if target >= row[0] && target <= row[n-1] {
			return binarySearch(row, target)
		}
	}
	return false
}

func binarySearch(row []int, target int) bool {
	if len(row) == 0 {
		return false
	}
	low, high := 0, len(row)-1
	for low <= high {
		mid := (low + high) / 2
		if row[mid] == target {
			return true
		} else if target < row[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

// func main() {
// 	matrix := [][]int{{1, 3}}
// 	fmt.Println(searchMatrix(matrix, 3))
// }
