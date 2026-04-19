func singleNumber(nums []int) []int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
        if m[num] == 2 {
            delete(m, num)
        }
    }

    result := make([]int, 2)
    index := 0
    for key, _ := range m {
        result[index] = key
        index++
    }

    return result
}