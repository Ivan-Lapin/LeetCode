func minNumberOfFrogs(croakOfFrogs string) int {

	counts := make([]int, 5)
    active := 0
    maxActive := 0

	for i := 0; i < len(croakOfFrogs); i++ {
		switch croakOfFrogs[i] {
		case 'c':
			counts[0]++
            active++
            if active > maxActive {
                maxActive = active
            }
		case 'r':
            if counts[0] == 0 {
				return -1
			}
			counts[1]++
            counts[0]--
		case 'o':
            if counts[1] == 0 {
				return -1
			}
			counts[2]++
            counts[1]--
		case 'a':
            if counts[2] == 0 {
				return -1
			}
			counts[3]++
            counts[2]--
		default:
            if counts[3] == 0 {
				return -1
			}
            counts[3]--
            active--
		}
	}

    for i := range counts {
        if counts[i] != 0 {
            return -1
        }
    }

	return maxActive
}