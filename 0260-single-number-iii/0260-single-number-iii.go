func singleNumber(nums []int) []int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }

    result := make([]int, 2)
    index := 0
    for key, val := range m {
        if val == 1 {
            result[index] = key
            index++
        }
        if index == 2 {
            break
        }
    }

    return result
}