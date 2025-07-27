package main

func findMissingAndRepeatedValuesArray2(nums []int) []int {
	xr := 0
	for i, num := range nums {
		xr = xr ^ num
		xr = xr ^ (i + 1)
	}
	// find out the first set bit from right
	// so that we can create two different group of nos
	// for nos which have bits set at the same place and for nos which does not have nos set at the same place
	setBitFromRight := 0
	for {
		if xr&(1<<setBitFromRight) != 0 {
			break
		}
		setBitFromRight++
	}
	zero, one := 0, 0
	for _, num := range nums {
		if num&(1<<setBitFromRight) != 0 {
			one = one ^ num
		} else {
			zero = zero ^ num
		}
	}
	for i := 1; i <= len(nums); i++ {
		if i&(1<<setBitFromRight) != 0 {
			one = one ^ i
		} else {
			zero = zero ^ i
		}
	}
	count := 0
	for _, num := range nums {
		if num == zero {
			count++
		}
	}
	if count == 2 {
		return []int{zero, one}
	} else {
		return []int{one, zero}
	}
}

// func main() {
// 	nums := []int{1, 3, 4, 5, 3}
// 	ans := findMissingAndRepeatedValuesArray2(nums)
// 	fmt.Println(ans) // [3 2]
// }
