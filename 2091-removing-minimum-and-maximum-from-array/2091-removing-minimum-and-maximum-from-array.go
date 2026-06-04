func minimumDeletions(nums []int) int {

    maxVal := nums[0]
    leftMax := 1
    rightMax := 1

    minVal := nums[0]
    leftMin := 1
    rightMin := 1

    for i, num := range nums {
        if num >= maxVal {
            maxVal = num
            leftMax = i+1
            rightMax = len(nums)-i
        }

        if num <= minVal {
            minVal = num
            leftMin = i+1
            rightMin = len(nums)-i
        }
    }

    case1 := leftMax + rightMin
    case2 := leftMin + rightMax

    var case3, case4 int
    if leftMax >= leftMin {
        case3 = leftMax
    } else {
        case3 = leftMin
    }

    if rightMax >= rightMin {
        case4 = rightMax
    } else {
        case4 = rightMin
    }

    return min(case1, case2, case3, case4)
}