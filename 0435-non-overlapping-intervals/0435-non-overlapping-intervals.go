func eraseOverlapIntervals(intervals [][]int) int {

    if len(intervals) == 1 {
        return 0
    }

    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

    lastEnd := intervals[0][1]
    counts := 1
    
    for i := 1; i < len(intervals); i++{
        if intervals[i][0] >= lastEnd {
            lastEnd = intervals[i][1]
            counts++
        }
    }

    return len(intervals) - counts
	
}