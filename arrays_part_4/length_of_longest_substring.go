package main

/*
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

func lengthOfLongestSubstring_OOf2NTime(s string) int {
	sLen := len(s)
	sSlice := []rune(s)
	l, r := 0, 0
	ans := 0
	charMap := make(map[rune]struct{})
	var currChar rune
	var currCharExists bool

	for r < sLen {
		currChar = sSlice[r]
		for {
			_, currCharExists = charMap[currChar]
			if !currCharExists {
				break
			}
			delete(charMap, sSlice[l])
			l++
		}
		charMap[currChar] = struct{}{}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func lengthOfLongestSubstring_OOfNTime(s string) int {
	sLen := len(s)
	sSlice := []rune(s)
	l, r := 0, 0
	ans := 0
	charMap := make(map[rune]int)
	var currChar rune
	var currCharExists bool
	var existingIndexOfCurrChar int

	for r < sLen {
		currChar = sSlice[r]
		if existingIndexOfCurrChar, currCharExists = charMap[currChar]; currCharExists {
			if existingIndexOfCurrChar >= l {
				l = existingIndexOfCurrChar + 1
			}
		}
		charMap[currChar] = r
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
