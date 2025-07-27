package main

import "sort"

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
*/

func merge(intervals [][]int) [][]int {
	// sort intervals based on first element
	// then based on 2nd element
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := [][]int{}

	// Algo
	// 1. iterate for each interval in intervals
	// 2. check if current inverval is already part of last interval in res
	//  if yes then skip looking to merge further
	// 3. if current interval is not part of last ans interval
	// then check in intervals ahead if they can be merged with this interval
	// if yes then take end as max(end, currEnd)
	// else break
	// 4. append solution to answer as []int{start, end}
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		start := interval[0]
		end := interval[1]
		if len(res) > 0 && end <= res[len(res)-1][1] {
			continue
		}
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] <= end {
				if intervals[j][1] > end {
					end = intervals[j][1]
				}
			} else {
				break
			}
		}
		res = append(res, []int{start, end})
	}
	return res
}
