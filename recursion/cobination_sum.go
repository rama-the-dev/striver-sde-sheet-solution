package main

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}

	var dfs func(index, currTarget int, currCombination []int)

	dfs = func(currIndex, currTarget int, currCombination []int) {
		if currIndex == len(candidates) {
			if currTarget == 0 {
				currCombinationCopy := make([]int, len(currCombination))
				copy(currCombinationCopy, currCombination)
				ans = append(ans, currCombinationCopy)
			}
			return
		}

		// include current no and do next recursion with same no again
		if candidates[currIndex] <= currTarget {
			currCombination = append(currCombination, candidates[currIndex])
			dfs(currIndex, currTarget-candidates[currIndex], currCombination)
			currCombination = currCombination[:len(currCombination)-1]
		}

		// do not include current no and do next recursion with next index and element
		dfs(currIndex+1, currTarget, currCombination)
	}

	dfs(0, target, []int{})
	return ans
}

func combinationSumOptimized(candidates []int, target int) [][]int {
	ans := [][]int{}

	var dfs func(index, currTarget int, currCombination []int)

	dfs = func(currIndex, currTarget int, currCombination []int) {
		if currIndex == len(candidates) {
			return
		}
		if currTarget == 0 {
			currCombinationCopy := make([]int, len(currCombination))
			copy(currCombinationCopy, currCombination)
			ans = append(ans, currCombinationCopy)
			return
		}

		// include current no and do next recursion with same no again
		if candidates[currIndex] <= currTarget {
			currCombination = append(currCombination, candidates[currIndex])
			dfs(currIndex, currTarget-candidates[currIndex], currCombination)
			currCombination = currCombination[:len(currCombination)-1]
		}

		// do not include current no and do next recursion with next index and element
		if currTarget > 0 {
			dfs(currIndex+1, currTarget, currCombination)
		}
	}

	dfs(0, target, []int{})
	return ans
}

// func main() {
// 	nums := []int{2, 3, 6, 7}
// 	ans := combinationSumOptimized(nums, 7)
// 	fmt.Println(ans)
// }
