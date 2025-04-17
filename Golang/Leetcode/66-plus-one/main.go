package plus_one

import (
	"fmt"
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	var arrStr []string
	for _, v := range digits {
		arrStr = append(arrStr, strconv.Itoa(v))
	}
	str1 := ""
	for _, v := range arrStr {
		str1 = str1 + v
	}
	largeNumber := new(big.Int)
	num1, _ := largeNumber.SetString(str1, 10)
	largeNumber.Add(num1, big.NewInt(1))
	str2 := largeNumber.String()
	resultStr := make([]int, len(str2))
	for i, v := range str2 {
		resultStr[i] = int(v - '0')
	}

	fmt.Println(resultStr)
	return resultStr
}

func main() {
	plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
}

// my solution for simple numbers
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func plusOne(digits []int) []int {
// 	num := 0
// 	for i := 0; i < len(digits); i++ {
// 		pow := len(digits) - 1 - i
// 		num += digits[i] * int(math.Pow(10, float64(pow)))
// 	}

// 	resultNum := num + 1

// 	countDigits := int(math.Log10(float64(resultNum))) + 1
// 	result := make([]int, countDigits)

// 	for i := countDigits - 1; i >= 0; i-- {
// 		result[i] = resultNum % 10
// 		resultNum = resultNum / 10
// 	}

// 	return result
// }
