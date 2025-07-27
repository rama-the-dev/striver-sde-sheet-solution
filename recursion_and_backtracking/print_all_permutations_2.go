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
func permute2(nums []int) [][]int {
	// algo
	// we need recursion + backtracking
	// so in approach - 2
	// we swap numbers starting from current index to last index
	// and after each swap run recursion once again
	// and backtrack to get the next permutation

	fact := factorial2(len(nums))
	ans := make([][]int, fact)

	ansIndex := 0
	var recursion func(currIndex int, currPerm []int)
	recursion = func(currIndex int, currPerm []int) {
		if currIndex == len(nums) {
			currPermCopy := make([]int, len(nums))
			copy(currPermCopy, currPerm)
			ans[ansIndex] = currPermCopy
			ansIndex++
			return
		}
		// iterate over all the number starting from current index
		// swap numbers from current index to last index one by one and run recursion for each swap
		for i := currIndex; i < len(nums); i++ {
			nums[currIndex], nums[i] = nums[i], nums[currIndex]
			recursion(currIndex+1, nums)
			nums[currIndex], nums[i] = nums[i], nums[currIndex]
		}
	}
	recursion(0, nums)
	return ans
}

func factorial2(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	ans := permute2(nums)
// 	fmt.Println(ans)
// }
