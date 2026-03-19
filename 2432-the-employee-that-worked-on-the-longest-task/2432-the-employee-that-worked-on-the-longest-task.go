func hardestWorker(n int, logs [][]int) int {
    maxTime := 0
    id := 0
    start := 0
    for _, log := range logs {
        work := log[1] - start
        if work > maxTime {
            maxTime = work
            id = log[0] 
        } else if work == maxTime && log[0] < id {
            id = log[0]
        }
        start = log[1]
    }
    return id
}