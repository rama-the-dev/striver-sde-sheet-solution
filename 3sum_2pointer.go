package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var l, r int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r = i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r -= 1
			} else if sum < 0 {
				l += 1
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l += 1
				r -= 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
				for nums[r] == nums[r+1] && l < r {
					r -= 1
				}
			}
		}
	}
	return res
}

// func main() {
// 	nums := []int{-1, 0, 1, 2, -1, -4}
// 	fmt.Println(threeSum(nums))
// }
