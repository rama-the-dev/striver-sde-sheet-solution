package main

import "strconv"

/*
60. Permutation Sequence
Solved
Hard
Topics
premium lock icon
Companies
The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.



Example 1:

Input: n = 3, k = 3
Output: "213"
*/

func getPermutation(n int, k int) string {
	// algo
	// create array of numbers from 1-n
	// since arrays are 0 indexed, so it'll boil down to k-1th permutation
	// fill number one by from left to right
	// if we take one number out, then how many permutations we can generate with left out numbers
	// so using this logic, we can get the number and reduce problem to shorter version
	nums := make([]int, n)
	factOfNumsLenMinusOne := 1 // for n = 4, it'll calculate 3! = 6
	for i := 1; i < n; i++ {
		nums[i-1] = i
		factOfNumsLenMinusOne *= i
	}
	nums[n-1] = n
	k -= 1 // reduce one as array is 0 indexed

	ans := ""
	var indexOfCurrentNum int
	for {
		indexOfCurrentNum = k / factOfNumsLenMinusOne // 16/6 = 2
		ans += strconv.Itoa(nums[indexOfCurrentNum])  // nums[2] = 3
		// remove number taken from array into answer from nums
		nums = removeAtIndex(nums, indexOfCurrentNum)
		// if nums len is 0 then break from loop
		if len(nums) == 0 {
			break
		}
		// now problem is reduced, so n-1 numbers have to be filled up
		k %= factOfNumsLenMinusOne         // 16 % 6 = 4, so find 4th permutation in [1, 2, 4]
		factOfNumsLenMinusOne /= len(nums) // 6 / 3 = 2, so after one number taken, group of 2 will be possible to form
	}
	return ans
}

func removeAtIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
