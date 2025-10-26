func numberOfArithmeticSlices(nums []int) int {

    if len(nums) < 3 {
        return 0
    }

    count := 0
    left := 0
    val := nums[1] - nums[0]
    
    for l, r := 1, 2; r < len(nums) && r > l; {

        if  nums[r] - nums[l] != val || r == len(nums) - 1 {

            k := 0
            if nums[r] - nums[l] != val {
                k = r - left 
            } else {
                k = r - left + 1
            }

            if k >= 3 {
                count += (k-1)*(k-2)/2
            }

            val = nums[r] - nums[l]
            left = l
        }

        if r != len(nums) - 1 {
            r++
        }
        l++

    }

    return count
}