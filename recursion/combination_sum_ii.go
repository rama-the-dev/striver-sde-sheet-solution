package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	// algo
	// run recursion for each no including
	// and backtrack and run without that no
	// exit conditions of recursion
	// when currIndex == len(candidates)
	// and in case case currTarget is 0 then currCombination becomes answer
	// to ensure we do not have duplicates
	// sort candidates
	// and while running recursion make sure when next recursion going without
	// current no then next no should not be equal to current no
	// so that we can ensure no duplicity
	sort.Ints(candidates)
	ans := [][]int{}
	var dfs func(currIndex, currTarget int, currCombination []int)
	dfs = func(currIndex, currTarget int, currCombination []int) {
		if currIndex == len(candidates) {
			if currTarget == 0 {
				currCombinationCopy := make([]int, len(currCombination))
				copy(currCombinationCopy, currCombination)
				ans = append(ans, currCombinationCopy)
			}
			return
		}
		// include current no
		// if currNum <= currTarget
		// run recursion and then backtrack that is remove current number
		currNum := candidates[currIndex]
		if currNum <= currTarget {
			currCombination = append(currCombination, currNum)
			dfs(currIndex+1, currTarget-currNum, currCombination)
			currCombination = currCombination[:len(currCombination)-1]
		}

		for currIndex+1 < len(candidates) && candidates[currIndex+1] == candidates[currIndex] {
			currIndex++
		}
		dfs(currIndex+1, currTarget, currCombination)
	}

	dfs(0, target, []int{})
	return ans
}

// func main() {
// 	nums := []int{2, 1, 2}
// 	ans := combinationSum2(nums, 2)
// 	fmt.Println(ans)
// }
