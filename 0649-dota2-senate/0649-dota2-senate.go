func predictPartyVictory(senate string) string {

	runes := []rune(senate)

	countR := 0
	countD := 0
	countBan := 0

	for i := 0; i < len(runes)*5 || countR == countD; i++ {
		index := i % len(runes)
		if runes[index] == 0 {
			continue
		}
		if runes[index] == 'R' {
			if countD > 0 {
				runes[index] = 0
				countBan++
				countD--
			} else {
				countR++
			}
		}
		if runes[index] == 'D' {
			if countR > 0 {
				runes[index] = 0
				countBan++
				countR--
			} else {
				countD++
			}
		}
	}

	if countR > countD {
		return "Radiant"
	} else {
		return "Dire"
	}
}