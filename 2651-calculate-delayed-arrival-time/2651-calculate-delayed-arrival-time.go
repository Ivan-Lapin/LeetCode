func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
    for range delayedTime {
        arrivalTime += 1
        if arrivalTime == 24 {
            arrivalTime = 0
        }
    }
    return arrivalTime
}