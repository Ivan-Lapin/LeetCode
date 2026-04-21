func nextGreaterElements(nums []int) []int {

	m := make(map[int][]int)
	set := make(map[int]bool)

	for i := range nums {
		set[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; ; j++ {
			if j == len(nums) {
				j = 0
			}

			if j == i {
				break
			}
			if nums[i] < nums[j] {
				m[nums[i]] = append(m[nums[i]], nums[j])
				break
			}
		}
	}

	ans := make([]int, len(nums))
	for i := range len(nums) {
		if _, exist := m[nums[i]]; !exist || len(m[nums[i]]) == 0 {
			ans[i] = -1
		} else {
			ans[i] = m[nums[i]][0]
			m[nums[i]] = m[nums[i]][1:]
		}
	}

	return ans
}