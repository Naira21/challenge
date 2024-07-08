package main

func applyOperations(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums[i-1] = nums[i] * 2
			nums[i] = 0
		}
	}
	//переміщуємо всі ненульові значення на початок масиву
	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			if lastNonZeroFoundAt != i {
				nums[i] = 0
			}
			lastNonZeroFoundAt++
		}
	}
	return nums
}

func main() {
	applyOperations([]int{1, 2, 2, 1, 1, 0})

}

// https://leetcode.com/problems/apply-operations-to-an-array/
