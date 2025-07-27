package main

import "fmt"

func uniquePaths(m int, n int) int {
	var recursion func(currI, currJ int) int
	recursion = func(currI, currJ int) int {
		if currI >= m || currJ >= n {
			return 0
		} else if currI == m-1 && currJ == n-1 {
			return 1
		}

		ans := 0
		// try right move
		ans += recursion(currI+1, currJ)
		// try bottom move
		ans += recursion(currI, currJ+1)

		// return ans
		return ans
	}
	return recursion(0, 0)
}

func main() {
	ans := uniquePaths(3, 7)
	fmt.Println(ans)
}
