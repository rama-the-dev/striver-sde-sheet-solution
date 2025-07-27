package main

/*
18. 4Sum
Solved
Medium
Topics
Companies
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
*/

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			st := make(map[int]struct{})
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				fourth := target - sum
				if _, exists := st[fourth]; exists {
					sol := []int{nums[i], nums[j], nums[k], fourth}
					sort.Ints(sol)
					if !isDuplicate(res, sol) {
						res = append(res, sol)
					}
				}
				st[nums[k]] = struct{}{}
			}
		}
	}
	return res
}

func isDuplicate(res [][]int, sol []int) bool {
	for _, oneOfTheSol := range res {
		if isArrayEqual(oneOfTheSol, sol) {
			return true
		}
	}
	return false
}

func isArrayEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != 4 || len(arr2) != 4 {
		return false
	}
	for i, num := range arr1 {
		if num != arr2[i] {
			return false
		}
	}
	return true
}
