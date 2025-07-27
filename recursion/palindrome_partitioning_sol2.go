package main

import "fmt"

func partitionSol2(s string) [][]string {
	ans := [][]string{}

	var tryEnds func(start, end int, path []string)
	tryEnds = func(start, end int, path []string) {
		if end == len(s) {
			if start == len(s) {
				copyPath := make([]string, len(path))
				copy(copyPath, path)
				ans = append(ans, copyPath)
			}
			return
		}

		curr := s[start : end+1]
		if isPalindromeSol2(curr) {
			path = append(path, curr)
			tryEnds(end+1, end+1, path) // next substring starts from end+1
			path = path[:len(path)-1]   // backtrack
		}
		tryEnds(start, end+1, path) // simulate next iteration of for loop
	}

	tryEnds(0, 0, []string{})
	return ans
}

func isPalindromeSol2(s string) bool {
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

func main() {
	ans := partitionSol2("aab")
	fmt.Println(ans)
}
