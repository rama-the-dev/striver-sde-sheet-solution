package main

import "sort"

func merge1(intervals [][]int) [][]int {
	// sort intervals based on first element
	// then based on 2nd element
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := [][]int{}

	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if len(res) == 0 || (len(res) > 0 && interval[0] > res[len(res)-1][1]) {
			res = append(res, interval)
		} else {
			resLast := res[len(res)-1]
			resLast[1] = max(interval[1], resLast[1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge_2(intervals [][]int) [][]int {
	// sort matrix based on first element of rows
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	// now start iterating intervals one by one
	for _, interval := range intervals {
		// if ans is not empty
		if len(ans) != 0 {
			// take the last ans
			// and check if current interval can be merged with last ans
			lastAns := ans[len(ans)-1]
			if interval[0] <= lastAns[1] {
				lastAns[1] = max(lastAns[1], interval[1])
			} else {
				ans = append(ans, interval)
			}
		} else {
			ans = append(ans, interval)
		}
	}
	return ans
}
