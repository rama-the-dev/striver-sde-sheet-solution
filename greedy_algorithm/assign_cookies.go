package main

import (
	"fmt"
	"sort"
)

/*
Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.

Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.

Example 1:

Input: g = [1,2,3], s = [1,1]
Output: 1
Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
You need to output 1.

*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	// l to track childrens
	// r to track cookie
	l, r := 0, 0
	// iterate over cookies
	for r < len(s) && l < len(g) {
		// in case current cookie can satisfy current child
		// then assign else move to next cookie
		if s[r] >= g[l] {
			l++
		}
		r++
	}
	return l
}

func main() {
	ans := findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})
	fmt.Println(ans)
}
