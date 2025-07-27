package main

import "math"

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
*/

func majorityElement(nums []int) []int {
	// take 2 answers and their count
	// when count becomes zero check if current no not equal to another answer
	// if not then increase current count and
	ans1, ans2, count1, count2 := math.MinInt, math.MinInt, 0, 0
	for _, num := range nums {
		if count1 == 0 && num != ans2 {
			ans1, count1 = num, 1
		} else if count2 == 0 && num != ans1 {
			ans2, count2 = num, 1
		} else if num == ans1 {
			count1++
		} else if num == ans2 {
			count2++
		} else {
			count1--
			count2--
		}
	}
	reqCount := len(nums)/3 + 1
	c1, c2 := 0, 0
	for _, num := range nums {
		if num == ans1 {
			c1++
		} else if num == ans2 {
			c2++
		}
	}
	ans := []int{}
	if c1 >= reqCount {
		ans = append(ans, ans1)
	}
	if c2 >= reqCount {
		ans = append(ans, ans2)
	}
	return ans
}
