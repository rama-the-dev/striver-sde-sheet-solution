package main

import "sort"

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and using only constant extra space.



Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
*/

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		currNum, prevNum := nums[i], nums[i-1]
		if currNum == prevNum {
			return currNum
		}
	}
	return 0
}
