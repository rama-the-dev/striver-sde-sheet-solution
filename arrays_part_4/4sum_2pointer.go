package main

import "sort"

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	res := [][]int{}

	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsLen; j++ {
			if j > i+1 && j < numsLen && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, numsLen-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum > target {
					l -= 1
				} else if sum < target {
					k += 1
				} else {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					l -= 1
					k += 1
					for k < l && nums[l] == nums[l+1] {
						l -= 1
					}
					for k < l && nums[k] == nums[k-1] {
						k += 1
					}
				}
			}
		}
	}
	return res
}
