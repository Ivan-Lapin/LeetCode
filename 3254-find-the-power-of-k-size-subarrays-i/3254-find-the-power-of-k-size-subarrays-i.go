func resultsArray(nums []int, k int) []int {

    if k == 1{
        return nums
    }

    result := make([]int, 0, len(nums)-k+1)
    strike := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] + 1{
            strike++
        } else {
            strike = 1
        }

        if i >= k-1 {
            if strike >= k{
                result = append(result, nums[i])
            } else {
                result = append(result, -1)
            }
        }
    }

    return result
}
