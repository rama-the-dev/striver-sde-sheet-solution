package main

func longestSubarraySumEqualsK_1(nums []int, k int) int {
	ans := 0
	l, r := 0, 0
	numsLen := len(nums)
	currSum := 0
	for l < numsLen && r < numsLen {
		for l <= r && currSum > k {
			currSum -= nums[l]
			l++
		}
		if currSum == k {
			ans = max(ans, r-l)
		}
		if r < numsLen {
			currSum += nums[r]
		}
		r++
	}
	return ans
}

// func main() {
// 	nums := []int{1, 2, 3, 1, 1, 0, 0, 1, 1, 1, 5}
// 	ans := longestSubarraySumEqualsK_1(nums, 5)
// 	fmt.Println(ans)
// }
