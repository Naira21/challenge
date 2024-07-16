package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)
	splitString := strings.Split(trimmed, " ")
	fmt.Println(len(splitString[len(splitString)-1]))
	return len(splitString[len(splitString)-1])
}

func main() {
	lengthOfLastWord("luffy is still joyboy")
}
