package main

func longestConsecutiveUsingSet(nums []int) int {
	// Algo
	// store numbers in hash map
	// for each number check
	// if number less than one of current number exist in map
	// it means current number is not the start of series
	// else
	// run a for loop and keep increasing one by one
	// and check upto what point numbers exist
	// and store total series length as ans if current series is of max length
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}
	maxSeriesLen := 0
	// now for each no check if no less than one exist
	for num := range numsMap {
		if _, ok := numsMap[num-1]; !ok {
			seriesLen := getSeriesLen(numsMap, num)
			maxSeriesLen = max(maxSeriesLen, seriesLen)
		}
	}
	return maxSeriesLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSeriesLen(numsMap map[int]struct{}, num int) int {
	seriesLen := 1
	delete(numsMap, num)
	nextNum := num + 1
	for {
		if _, ok := numsMap[nextNum]; ok {
			delete(numsMap, nextNum)
			seriesLen += 1
			nextNum += 1
		} else {
			break
		}
	}
	return seriesLen
}
