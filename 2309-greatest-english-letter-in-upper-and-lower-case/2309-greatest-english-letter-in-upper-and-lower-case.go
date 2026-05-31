func greatestLetter(s string) string {

	m := make(map[byte]bool)
	var result byte

	dif := 'a' - 'A'
	for i := range len(s) {
		m[s[i]] = true
		if s[i] >= 'a' {
			m[s[i]] = true
			if m[s[i]-byte(dif)] && s[i]-byte(dif) > result {
				result = s[i] - byte(dif)
			}
		} else {
			if m[s[i]+byte(dif)] && s[i] > result {
				result = s[i]
			}
		}
	}

	if result == 0 {
		return ""
	}

	return string(result)
}