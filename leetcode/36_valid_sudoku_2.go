package main

func isValidSudok2(board [][]byte) bool {
	// Algo
	// solve it using hasmaps for each row, column and box
	// create 9 maps for 9 rows to store whther specific digit is present
	rowHashMaps := make([]map[byte]bool, 9)
	colHashMaps := make([]map[byte]bool, 9)
	boxHashMaps := make([]map[byte]bool, 9)

	// initialize all hash maps internal maps
	for i := 0; i < 9; i++ {
		rowHashMaps[i] = make(map[byte]bool)
		colHashMaps[i] = make(map[byte]bool)
		boxHashMaps[i] = make(map[byte]bool)
	}

	// iterate through each element
	for i, row := range board {
		for j, elem := range row {
			// if current number is empty then no need to check anything
			if elem == '.' {
				continue
			}
			// else check if current element is not duplicate in row, col and box
			// 🧠 Rule of Thumb:
			// To convert 2D grid (r, c) to 1D array index:
			// index = r * numCols + c
			boxHashMapIndex := (i/3)*3 + (j / 3)
			// if current number is duplicate then return false
			if rowHashMaps[i][elem] || colHashMaps[j][elem] || boxHashMaps[boxHashMapIndex][elem] {
				return false
			}
			// else put current in all 3 hashmaps
			rowHashMaps[i][elem] = true
			colHashMaps[j][elem] = true
			boxHashMaps[boxHashMapIndex][elem] = true
		}
	}
	return true
}
