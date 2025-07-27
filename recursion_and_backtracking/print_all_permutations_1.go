package main

/*
46. Permutations
Solved
Medium
Topics
premium lock icon
Companies
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	// algo
	// we need recursion + backtracking
	// so in approach - 1
	// what we do is to pick one number and mark that position as picked
	// now call recursion to pick next number
	// if number at current poisition is not picked already, then pick that number
	// exit condition : when current answer length is equal to nums length then add that to answer and exit

	// create answer slice of size length of total numbers
	fact := factorial(len(nums))
	ans := make([][]int, fact)

	ansIndex := 0
	var recursion func(currPerm []int, positionMarker map[int]bool)
	recursion = func(currPerm []int, positionMarker map[int]bool) {
		if len(currPerm) == len(nums) {
			currPermCopy := make([]int, len(nums))
			copy(currPermCopy, currPerm)
			ans[ansIndex] = currPermCopy
			ansIndex++
			return
		}
		// iterate over all the number to pick next number
		// check if any of the next number is not picked
		// then pick that number, run recursion on that and backtrack
		for i, num := range nums {
			if !positionMarker[i] {
				currPerm = append(currPerm, num)
				positionMarker[i] = true
				recursion(currPerm, positionMarker)
				// backtrack
				currPerm = currPerm[:len(currPerm)-1]
				positionMarker[i] = false
			}
		}
	}
	positionMarker := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		positionMarker[i] = false
	}
	recursion([]int{}, positionMarker)
	return ans
}

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}
