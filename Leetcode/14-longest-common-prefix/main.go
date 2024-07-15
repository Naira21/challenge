package main

import (
	"fmt"
)

func main() {
	longestCommonPrefix([]string{"a"})
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	firstWord := strs[0]
	var prefix []byte
	for i := 0; i < len(firstWord); i++ {
		letter := firstWord[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != letter {
				return string(prefix)
			}
		}
		prefix = append(prefix, letter)
	}
	fmt.Println(string(prefix))
	return string(prefix)
}

//my solution (not passed all test during sumbitting time)

// if len(strs) == 0 {
// 	return ""
// }
// if len(strs) == 1 {
// 	fmt.Println(strs[0])
// 	return strs[0]
// }

// comparisonTo := strs[0]
// matchedLetter := []string{}
// firstWordSplitted := strings.Split(comparisonTo, "")
// for i := 1; i <= len(strs)-1; i++ {
// 	splittedWord := strings.Split(strs[i], "")
// 	if firstWordSplitted[i-1] == splittedWord[i-1] {
// 		matchedLetter = append(matchedLetter, firstWordSplitted[i-1])
// 	}
// 	fmt.Println(matchedLetter)
// }

// if len(matchedLetter) > 0 {
// 	return strings.Join(matchedLetter, "")
// }
// return ""
// https://leetcode.com/problems/longest-common-prefix
