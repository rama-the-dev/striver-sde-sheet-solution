package main

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2
*/

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	currMax := 0
	for i, num := range nums {
		if num == 1 {
			if i == 0 {
				ans = 1
				currMax = 1
			} else {
				currMax++
				ans = max(ans, currMax)
			}
		} else {
			currMax = 0
		}
	}
	return ans
}
