func diagonalPrime(nums [][]int) int {
    maxVal := 0
    for i := 0; i < len(nums); i++{
        temp := nums[i][i]
        if temp > maxVal && big.NewInt(int64(temp)).ProbablyPrime(0) {
            maxVal = nums[i][i]
        }

        temp = nums[i][len(nums[i])-i-1]
        if temp > maxVal && big.NewInt(int64(temp)).ProbablyPrime(0){
            maxVal = nums[i][len(nums[i])-i-1]
        }
    } 

    return maxVal
}