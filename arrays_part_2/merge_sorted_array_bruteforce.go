package main

func mergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	/* Solution 1
	   nums1 = nums1[:m]
	   nums1 = append(nums1, nums2...)
	   sort.Ints(nums1)
	*/

	/* Solution 2 */
	index := 0
	nums3 := make([]int, m+n)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if nums1[p1] < nums2[p2] {
			nums3[index] = nums1[p1]
			p1++
		} else {
			nums3[index] = nums2[p2]
			p2++
		}
		index++
	}

	for p1 < m {
		nums3[index] = nums1[p1]
		p1++
		index++
	}

	for p2 < n {
		nums3[index] = nums2[p2]
		p2++
		index++
	}

	for i, num := range nums3 {
		nums1[i] = num
	}
}
