func sortByBits(arr []int) []int {

    sort.Slice(arr, func(i, j int) bool {
        c1 := bits.OnesCount(uint(arr[i]))
        c2 := bits.OnesCount(uint(arr[j]))
        if c1 == c2 {
            return arr[i] < arr[j]
        }
        return c1 < c2
    })

    return arr
}