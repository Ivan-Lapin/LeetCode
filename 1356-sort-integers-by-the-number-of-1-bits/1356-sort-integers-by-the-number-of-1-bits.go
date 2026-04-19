func sortByBits(arr []int) []int {

    type Number struct {
        num int
        countBits int
    }

    counts := make([]Number, len(arr))
    var x int
    var count int

    for i, num := range arr {
        x = num
        for x > 0 {
            if x & 1 == 1{
                count++
            }
            x>>=1
        }
        counts[i] = Number{num: num, countBits: count}
        count = 0
    }

    sort.Slice(counts, func(i, j int) bool {
        if counts[i].countBits == counts[j].countBits {
            return counts[i].num < counts[j].num
        }
		return counts[i].countBits < counts[j].countBits
	})

    result := make([]int, len(counts))
    for i, number := range counts {
        result[i] = number.num
    }

    return result
}