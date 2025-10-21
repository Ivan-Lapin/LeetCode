func countPairs(deliciousness []int) int {
    const mod = 1_000_000_007
    freq := make(map[int64]int)
    count := 0

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

    for _, xv := range deliciousness {
        x := int64(xv)
        for _, target := range powers {
            y := target - x
            if c, ok := freq[y]; ok {
                count += c
                if count >= mod {
                    count -= mod
                }
            }
        }
        freq[x]++
    }
    return count
}
