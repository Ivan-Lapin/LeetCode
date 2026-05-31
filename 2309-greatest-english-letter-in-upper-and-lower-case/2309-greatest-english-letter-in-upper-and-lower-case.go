func greatestLetter(s string) string {

	m := make(map[byte]bool)
	var result byte

	for i := range len(s) {
		m[s[i]] = true
		if s[i] >= 97 {
			m[s[i]] = true
            if m[s[i]-32] && s[i]-32 > result{
				result = s[i]-32
			}
		} else {
			if m[s[i]+32] && s[i] > result{
				result = s[i]
			}
		}
	}

    if result == 0 {
		return ""
	}

	return string(result)
}