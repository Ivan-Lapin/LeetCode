func applyOperations(nums []int) []int {

    result := make([]int, len(nums))
    index := 0

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] != 0 {
            if nums[i] == nums[i+1] {
                result[index] = nums[i] * 2
                nums[i+1] = 0
            } else {
                result[index] = nums[i]
            }
            index++
        }
    }

    if nums[len(nums)-1] != 0 {
        result[index] = nums[len(nums)-1]
    }

    return result
}