package main

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
	return slow + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ans := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(ans)
}
