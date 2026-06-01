func findRightInterval(intervals [][]int) []int {

	m := make(map[int]int)
	for i := range intervals {
		m[intervals[i][0]] = i
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([]int, len(intervals))

	for i := range result {
		result[i] = -1
	}

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] == intervals[i][1] {
			result[m[intervals[i][0]]] = m[intervals[i][0]]
			continue
		}
		idx := sort.Search(len(intervals), func(j int) bool {
			return intervals[i][1] <= intervals[j][0]
		})
		if idx < len(intervals) {
			result[m[intervals[i][0]]] = m[intervals[idx][0]]
		}
	}

	return result
}