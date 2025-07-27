package main

func trap(heights []int) int {
	heightsLen := len(heights)
	if heightsLen == 1 || heightsLen == 2 {
		return 0
	}
	// store suffix max from right to left
	suffixMax := make([]int, heightsLen)
	suffixMax[heightsLen-1] = heights[heightsLen-1]
	for i := heightsLen - 2; i >= 0; i-- {
		suffixMax[i] = max(heights[i], suffixMax[i+1])
	}
	prefixMax := 0
	ans := 0
	// algo take the min height of prefix and suffix
	// and substract current height from it
	for i, height := range heights {
		prefixMax = max(prefixMax, height)
		ans += min(prefixMax, suffixMax[i]) - height
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
