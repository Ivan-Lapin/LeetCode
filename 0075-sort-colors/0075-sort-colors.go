func sortColors(nums []int)  {

    count0 := 0
    count1 := 0
    count2 := 0

    for _, num := range nums {
        if num == 0 {
            count0++
        } else if num == 1{
            count1++
        } else {
            count2++
        }
    }

    for i := range len(nums) {
        if count0 > 0 {
            nums[i] = 0
            count0--
            continue
        }
        if count1 > 0 {
            nums[i] = 1
            count1--
            continue
        }
        if count2 > 0 {
            nums[i] = 2
            count2--
        }
    }
}