package dontgivemefive

func dontGiveMeFive(start int, end int) int {
	var fiveStock []int
	for i := start; i <= end; i++ {
		if i%5 == 0 && i%10 != 0 {
			fiveStock = append(fiveStock, i)
		}
	}
	return end - start - len(fiveStock) + 1
}

// https://www.codewars.com/kata/5813d19765d81c592200001a/train/go

//other solutions
//1
// func countNumbersWithout5(start, end int) int {
// 	count := 0
// 	for i := start; i <= end; i++ {
// 		if !hasDigit5(i) {
// 			count++
// 		}
// 	}
// 	return count
// }
// func hasDigit5(num int) bool {
// 	if num < 0 {
// 		num = -num // make num positive if it's negative
// 	}
// 	for num > 0 {
// 		if num%10 == 5 {
// 			return true
// 		}
// 		num /= 10
// 	}
// 	return false
// }
// func main() {
// 	start := -4
// 	end := 17
// 	result := countNumbersWithout5(start, end)
// 	fmt.Printf("Count of numbers without 5 between %d and %d is: %d\n", start, end, result)
// }

//2
// func DontGiveMeFive(start int, end int) int {
// 	result := 0
// 	for i := start; i <= end; i++ {
// 		num := strconv.Itoa(i)
// 		if !strings.Contains(num, "5") {
// 			result += 1
// 		}
// 	}
// 	return result
// }
//3
// func DontGiveMeFive(start int, end int) (count int) {
// 	for i := start; i <= end; i++ {
// 		if strings.Contains(strconv.Itoa(i), "5") {
// 			continue
// 		}
// 		count++
// 	}
// 	return
// }
