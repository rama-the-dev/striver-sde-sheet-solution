package main

import (
	"sort"
)

func subsetsWithDupSol2(nums []int) [][]int {
	// algo
	sort.Ints(nums)
	// for each index iterate
	ans := [][]int{}

	var dfs func(i int, currArr []int)

	dfs = func(start int, currArr []int) {
		if start == len(nums) {
			currArrCopy := make([]int, len(currArr))
			copy(currArrCopy, currArr)
			ans = append(ans, currArrCopy)
			return
		}

		// include current element into current array
		currArr = append(currArr, nums[start])
		// and call recursion with next index and current num included
		dfs(start+1, currArr)

		// do not include current element into current array
		currArr = currArr[:len(currArr)-1]
		// and call recursion with next index and current num not included
		// but before doing that make sure next start index num is not same as current index num
		// else it'll result into duplicate output
		for start+1 < len(nums) && nums[start+1] == nums[start] {
			start++
		}
		dfs(start+1, currArr)
	}

	dfs(0, []int{})

	return ans
}

// func main() {
// 	nums := []int{1, 2, 2}
// 	ans := subsetsWithDupSol2(nums)
// 	fmt.Println(ans)
// }
