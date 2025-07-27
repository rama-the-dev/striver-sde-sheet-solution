package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// Algo
	// using DP for bottom up approach
	// at each character check from character upto max word length
	// and check if word can be formed
	// if yes check if upto index prior to the last character word was formed
	// if yes then only word break is possible
	// https://www.youtube.com/watch?v=hK6Git1o42c
	sLen := len(s)
	dp := make([]bool, sLen+1)
	maxWordLen := maxWordLenInDict(wordDict)
	wordMap := wordDictMap(wordDict)
	// put true at the first position to satisfy base condition
	// as blank string would yield true
	dp[0] = true
	var currWord string
	for i := 1; i <= sLen; i++ {
		for j := i - 1; j >= 0; j-- {
			currWord = s[j:i]
			// check if current word is not more than maxDictionaryWordLen
			// if yes then no need to check availability in dictionary
			if len(currWord) > maxWordLen {
				break
			}
			// now check if current word exist in dictionary
			// also check if string ending just before current word start
			// was present in dictionary
			// then only mark current word found in dict as true in dp
			if _, ok := wordMap[currWord]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLen]
}

func maxWordLenInDict(dict []string) int {
	maxLen := 0
	wordLen := 0
	for _, word := range dict {
		wordLen = len(word)
		if wordLen > maxLen {
			maxLen = wordLen
		}
	}
	return maxLen
}

func wordDictMap(wordDict []string) map[string]struct{} {
	wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	return wordMap
}

func main() {
	ans := wordBreak("applepenapple", []string{"apple", "pen"})
	fmt.Println(ans)
}
