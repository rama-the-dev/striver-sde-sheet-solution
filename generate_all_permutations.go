package main

import (
	"fmt"
	"strconv"
)

func permuteNumber(n int) []int {
	digits := []byte(strconv.Itoa(n)) // convert number to digit slice (as byte for easy swapping)
	var result []int

	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(digits) {
			num, _ := strconv.Atoi(string(digits))
			result = append(result, num)
			return
		}
		seen := map[byte]bool{} // avoid duplicate digits if needed
		for i := start; i < len(digits); i++ {
			if seen[digits[i]] {
				continue
			}
			seen[digits[i]] = true
			digits[start], digits[i] = digits[i], digits[start] // swap
			backtrack(start + 1)
			digits[start], digits[i] = digits[i], digits[start] // backtrack
		}
	}

	backtrack(0)
	return result
}

func main() {
	num := 123
	perms := permuteNumber(num)
	fmt.Println(perms) // Output: [123 132 213 231 312 321]
}
