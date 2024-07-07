package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	fmt.Println(j + 1)

	return j + 1

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
