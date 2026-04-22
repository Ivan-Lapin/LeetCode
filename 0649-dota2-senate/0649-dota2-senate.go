func predictPartyVictory(senate string) string {

    n := len(senate)

	countR := make([]int, 0, n)
	countD := make([]int, 0, n)

    for i := range n {
        if senate[i] == 'R' {
            countR = append(countR, i)
        } else {
            countD = append(countD, i)
        }
    }

	for i := 0; len(countR) > 0 && len(countD) > 0; i++ {
		r := countR[0]
        d := countD[0]
        countR = countR[1:]
        countD = countD[1:]

        if r < d {
            countR = append(countR, i+n)
        } else {
            countD = append(countD, i+n)
        }
	}

	if len(countR) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}