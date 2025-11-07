func combinationSum4(nums []int, target int) int {

    sort.Ints(nums)

    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i <= target; i++{
        dp[i] = -1
    }

    var def func(t int) int

    def = func(t int) int {

        if t < 0 {
            return 0
        }

        if dp[t] != -1 {
            return dp[t]
        }

        total := 0
        for _, num := range nums {
            if num > t {
                break
            }  

            total += def(t-num)
        }

        dp[t] = total
        return total
    }

    return def(target)
}