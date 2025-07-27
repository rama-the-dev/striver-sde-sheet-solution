package main

func findMissingAndRepeatedValuesArray(nums []int) []int {
	n := len(nums)
	sn := (n * (n + 1)) / 2
	s2n := (n * (n + 1) * (2*n + 1)) / 6
	s, s2 := 0, 0
	for _, num := range nums {
		s += num
		s2 += num * num
	}

	val1 := s - sn
	val2 := s2 - s2n
	val2 = val2 / val1
	x := (val1 + val2) / 2
	y := x - val1

	return []int{x, y}
}

// func main() {
// 	nums := []int{1, 3, 4, 5, 3}
// 	ans := findMissingAndRepeatedValuesArray(nums)
// 	fmt.Println(ans) // [3 2]
// }
