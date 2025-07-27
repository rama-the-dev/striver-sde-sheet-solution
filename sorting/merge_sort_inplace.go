package main

import "fmt"

func mergeInPlace(nums []int, low, mid, high int) {
	left, right := low, mid+1
	temp := make([]int, 0, high-low+1)
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

	for j := low; j <= high; j++ {
		nums[j] = temp[j-low]
	}
}

func mergeSortInPlace(nums []int, low, high int) {
	if low == high {
		return
	}
	mid := (low + high) / 2
	mergeSortInPlace(nums, low, mid)
	mergeSortInPlace(nums, mid+1, high)
	mergeInPlace(nums, low, mid, high)
}

func main() {
	arr := []int{2, 1, -1, 3, 1, 8, 10, 6, -6}
	fmt.Println("Original array:", arr)

	low, high := 0, len(arr)-1
	mergeSortInPlace(arr, low, high)
	fmt.Println("Sorted array:  ", arr)
}
