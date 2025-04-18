package main

import "fmt"

func removeDuplicates(nums []int) int {
	// розвязок в якому вихідний nums змінюється
	if len(nums) == 0 {
		return 0
	}
	j := 0 // вказівник для унікальних елементів
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i] // переміщуємо новий унікальний елемент на позицію j
		}
	}
	fmt.Println(j + 1)

	return j + 1

	// розвязок в якому вихідний nums не змінюється
	// counter := []int{nums[0]}
	// if len(nums) > 1 {

	// 	for i := 1; i <= len(nums)-1; i++ {
	// 		if nums[i] != nums[i-1] {
	// 			counter = append(counter, nums[i])
	// 		}
	// 	}
	// } else {
	// 	return 1
	// }
	// fmt.Print(counter)

	// return len(counter)
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 3, 4, 5, 5}
	removeDuplicates(nums)
}
