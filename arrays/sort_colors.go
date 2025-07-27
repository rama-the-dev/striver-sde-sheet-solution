package main

// dutch national flag algorithm
func sortColors(nums []int) {
	// intitution behind algo
	// algo : Dutch National Flag algo
	// three pointer approach
	// low, mid, high pointer
	// from 0 -> low-1 -> all zeros
	// from low -> mid-1 -> all ones
	// from mid -> high -> unsorted nums
	// from high+1 -> n-1 -> all twos

	numsLen := len(nums)
	low, mid, high := 0, 0, numsLen-1
	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
