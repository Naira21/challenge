package plus_one

func PlusOne(digits []int) []int {
	var res int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i]
		if s+res <= 9 {
			digits[i] += res
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

func main2() {
	plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
}
