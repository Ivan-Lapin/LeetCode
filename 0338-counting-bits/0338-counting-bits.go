func countBits(n int) []int {
	arr := make([]int, n+1)
    var count int
    var num int

	for i := 0; i <= n; i++ {
        count = 0
        num = i
        for num > 0 {
            if num&1 == 1{
                count++
            }
            num>>=1
        }
        arr[i] = count
        count = 0
	}

	return arr
}