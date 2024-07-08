package applyoperationonarray

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}

	}
	return nums
}
