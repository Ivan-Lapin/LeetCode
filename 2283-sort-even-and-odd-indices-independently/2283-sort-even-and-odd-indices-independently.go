func sortEvenOdd(nums []int) []int {

    odd := []int{}
    even := []int{}
    result := make([]int, len(nums))

    for i, num := range nums {
        if i % 2 == 0{
            even = append(even, num)
        } else {
            odd = append(odd, num)
        }
    }

    sort.Ints(even)
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })

    for i, step := 0, 0; i < len(odd) && i < len(even); i++{
        result[i+step] = even[i]
        step++
        result[i+step] = odd[i]
    }

    if len(even) > len(odd) {
        result[len(result)-1] = even[len(even)-1]
    } else if len(even) < len(odd) {
        result[len(result)-1] = odd[len(odd)-1]
    }


    return result
}