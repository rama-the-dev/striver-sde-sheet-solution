package main

/*
560. Subarray Sum Equals K
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

*/

func subarraySum(nums []int, k int) int {
	ans := 0
	currSum := 0
	prefixSumMap := make(map[int]int)
	extraSum := 0
	for _, num := range nums {
		currSum += num
		extraSum = currSum - k
		// if current sum is equal to 1 then also add one to answer
		// another way to handle this corner case is to store (0,1) in prefixSumMap
		if currSum == k {
			ans += 1
		}
		// if extra sum exists then all the extra sum occurrences
		if sumOccurrences, extraSumExists := prefixSumMap[extraSum]; extraSumExists {
			ans += sumOccurrences
		}
		prefixSumMap[currSum] += 1
	}
	return ans
}

// func main() {
// 	nums := []int{1}
// 	ans := subarraySum(nums, 0)
// 	fmt.Println(ans)
// }
