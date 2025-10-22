func findNonMinOrMax(nums []int) int {
    minNum := nums[0]
    maxNum := nums[0]
    result := -1

    for i := 1; i < len(nums); i++{
        if nums[i] > maxNum {
            result = maxNum
            maxNum = nums[i]
        } else if nums[i] < minNum {
            result = minNum
            minNum = nums[i]
        } else {
            result = nums[i]
        }
    }

    if result == minNum || result == maxNum {
        return -1
    }

    return result
}