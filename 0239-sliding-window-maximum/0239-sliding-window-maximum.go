func maxSlidingWindow(nums []int, k int) []int {

	if k == 1 {
		return nums
	}

	deque := make([]int, 0, k)
	res := make([]int, len(nums)-k+1)

	for i := range k {
		if len(deque) == 0 || nums[deque[0]] < nums[i] {
			deque = []int{i}
		} else {
			for j := 0; j <= len(deque); j++ {
				if j == len(deque) {
					deque = append(deque, i)
				}
				if nums[i] >= nums[deque[j]] {
					deque[j] = i
					deque = deque[:j+1]
					break
				}
			}
		}
	}

	res[0] = nums[deque[0]]

	l := 1
	r := l + k - 1
	for r < len(nums) {
		if deque[0] < l {
			deque = deque[1:]
		}
		if len(deque) == 0 || nums[deque[0]] < nums[r] {
			deque = []int{r}
		} else {
			for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[r] {
				deque = deque[:len(deque)-1]
			}
			deque = append(deque, r)
		}
		res[l] = nums[deque[0]]
		l++
		r++
	}

	return res
}