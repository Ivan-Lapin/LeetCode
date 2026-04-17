func reorganizeString(s string) string {

	chars := []byte(s)
	slices.Sort(chars)

	s = string(chars)

	freqLet := make(map[int][]byte)

	var result strings.Builder

	index := 0
	count := 1
	for index < len(s) {
		for index != len(s)-1 && s[index] == s[index+1] {
			count++
			index++
		}
		freqLet[count] = append(freqLet[count], s[index])
		index++
		count = 1
	}

	miss := []byte{}
	num := len(s)
	var past byte
	for num > 0 {
		for i := range freqLet[num] {
			if result.Len() > 0 && past == freqLet[num][i] {
				miss = append(miss, past)
			} else {
				past = freqLet[num][i]
				result.WriteByte(freqLet[num][i])
				if len(miss) > 0 && miss[0] != past {
					result.WriteByte(miss[0])
                    past = miss[0]
					miss = miss[1:]
				}
			}
			freqLet[num-1] = append(freqLet[num-1], freqLet[num][i])
		}
		num--
	}

    if len(miss) != 0 {
		return ""
	}

	str := result.String()
	for i := 0; i < len(s)-1; i++ {
		if str[i] == str[i+1] {
			return ""
		}
	}

	return result.String()
}