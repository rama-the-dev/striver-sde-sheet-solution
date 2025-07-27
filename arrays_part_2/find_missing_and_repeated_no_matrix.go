package main

/*
You are given a 0-indexed 2D integer matrix grid of size n * n with values in the range [1, n2]. Each integer appears exactly once except a which appears twice and b which is missing. The task is to find the repeating and missing numbers a and b.

Return a 0-indexed integer array ans of size 2 where ans[0] equals to a and ans[1] equals to b.



Example 1:

Input: grid = [[1,3],[2,2]]
Output: [2,4]
Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4].
Example 2:

Input: grid = [[9,1,7],[8,9,2],[3,4,6]]
Output: [9,5]
Explanation: Number 9 is repeated and number 5 is missing so the answer is [9,5].

*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	m := len(grid)
	mem := make([]int, m*m+1)
	for _, row := range grid {
		for _, elem := range row {
			mem[elem] += 1
		}
	}
	ans := []int{0, 0}
	for i, elem := range mem {
		if i == 0 {
			continue
		}
		if elem == 0 {
			ans[1] = i
		} else if elem == 2 {
			ans[0] = i
		}
	}
	return ans
}
