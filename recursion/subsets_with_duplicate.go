package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// algo
	sort.Ints(nums)
	// for each index iterate
	ans := [][]int{}

	var dfs func(i int, currArr []int)

	dfs = func(start int, currArr []int) {
		// append a copy of the current subset into answer
		subset := make([]int, len(currArr))
		copy(subset, currArr)
		ans = append(ans, subset)
		fmt.Printf("Current stack details, start: %d currArr: %v\n", start, currArr)

		for i := start; i < len(nums); i++ {
			fmt.Printf("inside loop, start: %d currArr: %v, currNum: %d\n", start, currArr, nums[i])
			// skip duplicates
			// though it takes the first element
			// but at the same level it removes duplicate
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			currArr = append(currArr, nums[i])
			dfs(i+1, currArr)
			currArr = currArr[:len(currArr)-1]
		}
		fmt.Printf("Call finished for start: %d, currArr: %v\n", start, currArr)
	}

	dfs(0, []int{})

	return ans
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	ans := subsetsWithDup(nums)
// 	fmt.Println(ans)
// }
