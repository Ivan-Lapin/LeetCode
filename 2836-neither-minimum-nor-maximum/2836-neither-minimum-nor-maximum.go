func findNonMinOrMax(nums []int) int {
    minNum := nums[0]
    maxNum := nums[0]
    result := -1

    for i := 1; i < len(nums); i++{
        if nums[i] > maxNum {
            maxNum = nums[i]
        } 
        if nums[i] < minNum {
            minNum = nums[i]
        } 
    }

    for _, num := range nums {
        if (num != maxNum) && (num != minNum) {
            result = num
        } 
    }

    return result
}