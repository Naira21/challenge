//https://leetcode.com/problems/palindrome-number/

package main

func main() {
	isPalindrome(121)
}

func isPalindrome(x int) bool {
	var newInt int
	initX := x
	for x > 0 {
		lastDigit := x % 10
		newInt = newInt*10 + lastDigit
		x = x / 10
	}
	if newInt == initX {
		return true
	}
	return false
}

// reverse int in Go https://golangbyexample.com/reverse-number-golang/
