func nextGreaterElements(nums []int) []int {

	ans := make([]int, len(nums))
    for i := range nums {
        ans[i] = -1
    }
    stack := []int{}

    for i := 0; i < len(nums)*2; i++{
        idx := i%len(nums)
        num := nums[idx]
        for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
            ans[stack[len(stack)-1]] = num
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, idx)
    }

	return ans
}