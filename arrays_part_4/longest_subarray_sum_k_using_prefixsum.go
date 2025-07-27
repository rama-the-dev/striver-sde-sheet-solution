package main

func longestSubarraySumEqualsK(nums []int, k int) int {
	prefixSumMap := make(map[int]int)
	ans := 0
	currSum := 0
	extraSum := 0
	for i, num := range nums {
		currSum += num
		extraSum = currSum - k
		if currSum == k {
			ans = max(ans, i-0+1)
		} else if extraSumEndIndex, extraSumExists := prefixSumMap[extraSum]; extraSumExists {
			ans = max(ans, i-extraSumEndIndex)
		}
		if _, currSumAlreadyExists := prefixSumMap[currSum]; !currSumAlreadyExists {
			prefixSumMap[currSum] = i
		}
	}
	return ans
}

// func main() {
// 	nums := []int{1, 2, 3, 1, 1, 0, 0, 1, 1, 1, 5}
// 	ans := longestSubarraySumEqualsK(nums, 5)
// 	fmt.Println(ans)
// }
