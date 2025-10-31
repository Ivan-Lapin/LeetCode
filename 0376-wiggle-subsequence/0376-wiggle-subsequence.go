func wiggleMaxLength(nums []int) int {

    n := len(nums)
    if n < 2 {
        return len(nums)
    }

    up, down := 1, 1

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            up = down + 1
        } else if nums[i] < nums[i-1] {
            down = up + 1
        }
    }

    if up > down {
        return up
    }

    return down
    
}