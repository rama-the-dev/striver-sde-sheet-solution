package main

import "fmt"

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, low, high int) int {
	reverseCount := 0
	if low == high {
		return reverseCount
	}
	mid := (low + high) / 2
	reverseCount += mergeSort(nums, low, mid)
	reverseCount += mergeSort(nums, mid+1, high)
	reverseCount += merge(nums, low, mid, high)
	return reverseCount
}

func merge(nums []int, low, mid, high int) int {
	left, right := low, mid+1
	temp := make([]int, 0, high-low+1)
	reverseCount := countReverse(nums, low, mid, high)
	for left <= mid && right <= high {
		if nums[left] < nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}
	temp = append(temp, nums[left:mid+1]...)
	temp = append(temp, nums[right:high+1]...)

	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}
	return reverseCount
}

func countReverse(nums []int, low, mid, high int) int {
	count := 0
	right := mid + 1
	// left part of array have to be compared with right part
	// so take 2 pointers
	// right pointer for any left item would keep moving upto the point
	// 1. it doesn't cross the boundary
	// 2. nums[left] > 2*nums[right]
	for left := low; left <= mid; left++ {
		for right <= high && nums[left] > 2*nums[right] {
			right++
		}
		// right pointer start from mid+1
		// hence all elements upto right woul
		count += right - (mid + 1)
	}
	return count
}

func main() {
	arr := []int{2, 1, -1, 3, 1, 8, 10, 6, -6}
	fmt.Println("Original array:", arr)

	count := reversePairs(arr)
	fmt.Println("Sorted array:  ", arr)
	fmt.Println("reversePairs:  ", count)
}
