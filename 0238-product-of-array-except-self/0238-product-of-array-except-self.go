func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	temp := make([]int, len(nums))

	for i := range result {
		result[i] = 1
		temp[i] = 1
	}

	for i, j := 0, len(nums)-1; i < len(nums)-1 && j > 0; i, j = i+1, j-1 {
		result[i+1] *= nums[i] * result[i]
		temp[j-1] *= nums[j] * temp[j]
	}

	for i := range result {
		result[i] *= temp[i]
	}

	return result
}