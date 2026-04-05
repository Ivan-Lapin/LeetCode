func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums []int, l int, r int) {
    if l >= r {
        return
    }

    i, j := l, r
    pivot := nums[l + (r-l)/2]

    for i <= j {
        for nums[i] < pivot {
            i++
        }
        for nums[j] > pivot {
            j--
        }
        if i <= j {
            temp := nums[i]
            nums[i] = nums[j]
            nums[j] = temp
            i++
            j--
        }
    }

    if l < j {
        quickSort(nums, l, j)
    }

    if i < r {
        quickSort(nums, i, r)
    }
}


