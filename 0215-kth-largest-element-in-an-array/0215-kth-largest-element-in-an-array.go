func findKthLargest(nums []int, k int) int {

    m := make(map[int]int)
    max := nums[0]

    for _, num := range nums {
        m[num]++
        if num > max {
            max = num
        }
    }

    num := max
    for k > 0 {
        if val, exist := m[num]; exist {
            k -= val 
            if k <= 0 {
                return num
            }
        }
        num --
    }

    return num

}