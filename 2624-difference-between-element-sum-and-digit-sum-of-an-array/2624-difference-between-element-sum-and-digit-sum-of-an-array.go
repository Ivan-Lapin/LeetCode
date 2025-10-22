func differenceOfSum(nums []int) int {
    el_sum := 0
    digit_sum := 0

    for _, num := range nums {
        el_sum += num
        if num <= 9 {
            digit_sum += num
        } else {
            for num > 0 {
                s := num%10
                num /= 10
                digit_sum += s
            }
        }
    }

    return int(math.Abs(float64(el_sum - digit_sum)))
}