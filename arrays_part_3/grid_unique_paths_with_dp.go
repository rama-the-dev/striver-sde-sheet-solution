package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		row := make([]int, n)
		for i := range row {
			row[i] = -1
		}
		dp[i] = row
	}
	return countPaths(0, 0, m, n, dp)
}

func countPaths(i, j, m, n int, dp [][]int) int {
	if i >= m || j >= n {
		return 0
	} else if i == m-1 && j == n-1 {
		return 1
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	bottomTraversalPathCount := countPaths(i+1, j, m, n, dp)
	rightTraversalPathCount := countPaths(i, j+1, m, n, dp)
	dp[i][j] = bottomTraversalPathCount + rightTraversalPathCount
	return dp[i][j]
}
