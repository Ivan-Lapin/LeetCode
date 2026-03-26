func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	temp := make([]int, len(nums))

	for i := range result {
		result[i] = 1
		temp[i] = 1
	}

	for i := 0; i < len(nums)-1; i++ {
		result[i+1] *= nums[i] * result[i]
	}

	for i := len(nums) - 1; i > 0; i-- {
		temp[i-1] *= nums[i] * temp[i]
	}

	for i := range result {
		result[i] *= temp[i]
	}

	return result
}