package main

import "sort"

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max, temp := 1, 1
	lastNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNum {
			continue
		}
		if nums[i] == lastNum+1 {
			temp += 1
		} else {
			temp = 1
		}
		lastNum = nums[i]
		if temp > max {
			max = temp
		}
	}
	return max
}
