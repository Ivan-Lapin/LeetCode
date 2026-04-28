func subsetXORSum(nums []int) int {

    n := len(nums)
	totalXOR := 0
	
    for mask := 0; mask < (1 << n); mask++ {
        xor := 0

        for i := range n {
            if mask & (1 << i) != 0 {
                xor ^= nums[i]
            }
        }

        totalXOR += xor
    }

    return totalXOR

}