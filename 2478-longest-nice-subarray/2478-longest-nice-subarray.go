func longestNiceSubarray(nums []int) int {

    maxLen := 1
    lenght := 1
    mask := nums[0]

    for left, right := 0, 1; right < len(nums); right++{

        for mask & nums[right] != 0 {
            mask ^= nums[left]
            left++
        }

        mask |= nums[right]
        lenght = right - left + 1

        if maxLen < lenght {
            maxLen = lenght
        }

    }

    return maxLen
}