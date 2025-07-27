package main

func partition(s string) [][]string {
	ans := [][]string{}
	var dfs func(currIndex int, currAns []string)
	dfs = func(currIndex int, currAns []string) {
		if currIndex == len(s) {
			currAnsCopy := make([]string, len(currAns))
			copy(currAnsCopy, currAns)
			ans = append(ans, currAnsCopy)
			return
		}
		for i := currIndex; i < len(s); i++ {
			currSubString := s[currIndex : i+1]
			if isPalindrome(currSubString) {
				currAns = append(currAns, currSubString)
				dfs(i+1, currAns)
				currAns = currAns[:len(currAns)-1]
			}
		}
	}
	dfs(0, []string{})
	return ans
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// func main() {
// 	ans := partition("aab")
// 	fmt.Println(ans)
// }
