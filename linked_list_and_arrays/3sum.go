package main

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			currSum := nums[i] + nums[j] + nums[k]
			if currSum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if currSum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
