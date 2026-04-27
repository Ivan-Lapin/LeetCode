func readBinaryWatch(turnedOn int) []string {

	result := make([]string, 0, 30)

	if turnedOn > 8 {
		return result
	}

	count1 := 0
	count2 := 0
	var str string
	for i := 0; i <= 11; i++ {
		count1 = bits.OnesCount(uint(i))
		for j := 0; j <= 59; j++ {
			count2 = bits.OnesCount(uint(j))
			if count1+count2 == turnedOn {
				if j < 10 {
					str = fmt.Sprintf("%d:0%d", i, j)
				} else {
					str = fmt.Sprintf("%d:%d", i, j)
				}
				result = append(result, str)
			}
		}
	}

	return result
}