package main

import "math"

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
*/

func maxSubArray(nums []int) int {
	max := math.MinInt32
	sum := 0
	for _, num := range nums {
		sum += num

		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	return max
}
