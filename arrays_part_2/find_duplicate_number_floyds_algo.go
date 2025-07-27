package main

import "fmt"

func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}
}

func findDuplicate3(nums []int) int {
	// floyd warshall algo
	slow, fast := 0, 0
	// find the first meeting point in cycle
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	// move slow pointer at the start of array
	// now move both of them one by one and at next meeting point lies the duplicate
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	ans := findDuplicate3(nums)
	fmt.Println((ans))
}
