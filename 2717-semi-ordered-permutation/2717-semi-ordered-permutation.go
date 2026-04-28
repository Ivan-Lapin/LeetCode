func semiOrderedPermutation(nums []int) int {

	count := 0
	for i := 0; i < len(nums); {
		if nums[i] == 1 && nums[0] != 1{
            old_index := i
			for nums[0] != 1 {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				count++
				i--
			}
            i = old_index
		}

		if nums[i] == len(nums) && nums[len(nums)-1] != len(nums){
            old_index := i
			for nums[len(nums)-1] != len(nums) {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				count++
				i++
			}
            i = old_index
            continue
		}

        if nums[0] == 1 && nums[len(nums)-1] == len(nums) {
            break
        }

        i++
	}

	return count
}