func maxOperations(nums []int) int {

    count := 1
    sum := nums[0] + nums[1]

    for i := 2; i < len(nums) - 1; i += 2{
        if nums[i] + nums[i+1] == sum {
            count ++
        } else {
            break
        }
    }

    return count
}