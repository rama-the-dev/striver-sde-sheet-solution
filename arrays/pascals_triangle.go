package main

func generate(numRows int) [][]int {
	res := [][]int{[]int{1}}
	if numRows == 1 {
		return res
	}

	res = append(res, []int{1, 1})
	if numRows == 2 {
		return res
	}

	for i := 2; i < numRows; i++ {
		nextRow := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				nextRow[j] = 1
			} else {
				nextRow[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, nextRow)
	}

	return res
}
