func findClosest(x int, y int, z int) int {
    dist_1_3 := math.Abs(float64(z - x))
    dist_2_3 := math.Abs(float64(y - z))
    if dist_1_3 < dist_2_3 {
        return 1
    } else if dist_1_3 > dist_2_3 {
            return 2
    } else {
        return 0
    }
}