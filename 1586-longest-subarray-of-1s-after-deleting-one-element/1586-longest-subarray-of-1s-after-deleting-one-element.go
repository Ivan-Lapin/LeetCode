func longestSubarray(nums []int) int {

    maxLen := 0
    lenght := 0
    left := 0
    sum := 0

    flag := false

    for i, num := range nums {
        if num == 0 {
            if flag {
                if lenght > maxLen {
                    maxLen = lenght
                }
                lenght = len(nums[left:i]) - 1

            } else {
                flag = true
            }
            left = i
        } else {
            lenght++
            sum++
        }
    }

    if sum == len(nums) {
        return sum - 1
    }

    if lenght > maxLen {
        return lenght
    }

    return maxLen
}