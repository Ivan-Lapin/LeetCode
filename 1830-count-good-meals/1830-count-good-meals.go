func countPairs(deliciousness []int) int {
    const mod = 1_000_000_007
    count := 0

    freqMap := make(map[int64]int)

    maxVal := int64(0)
    for _, v := range deliciousness {
        if int64(v) > maxVal {
            maxVal = int64(v)
        }
    }
    maxSum := maxVal * 2

    powers := []int64{1}
    for p := int64(2); p <= maxSum; p <<= 1 {
        powers = append(powers, p)
    }

    for _, x := range deliciousness {
        for _, target := range powers {
            y := target - int64(x)
            if val, exist := freqMap[y]; exist {
                count += val
                if count >= mod {
                    count -= mod
                }
            }
        }
        freqMap[int64(x)]++
    }

    return count
}