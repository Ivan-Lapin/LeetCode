func longestOnes(nums []int, k int) int {

    flips := 0
    count := 0
    maxCount := 0
    nulls_index := []int{}
    null_index := 0

    for l, r := 0, 0; r < len(nums); r++{

        if nums[r] == 0 {

            nulls_index = append(nulls_index, r)
            flips++

            if flips > k {

                if count > maxCount {
                    maxCount = count
                }
                
                count -= len(nums[l:nulls_index[null_index]+1])
                l = nulls_index[null_index]+1
                flips--
                null_index++

            } 
        } 
        count++
    }

    if count > maxCount {
        return count
    }

    return maxCount
}