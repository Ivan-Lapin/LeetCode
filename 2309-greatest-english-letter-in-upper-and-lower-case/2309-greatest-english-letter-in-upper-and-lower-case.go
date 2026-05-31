func greatestLetter(s string) string {

	lower := [26]bool{}
	upper := [26]bool{}
	var result byte

	dif := 'a' - 'A'
	for i := range len(s) {

		if s[i] >= 'a' {
			idx := s[i] - 'a'
			lower[idx] = true
			if upper[idx] && idx >= result {
				result = s[i] - byte(dif)
			}
		} else {
			idx := s[i] - 'A'
			upper[idx] = true
			if lower[idx] && s[i] >= result {
				result = s[i]
			}
		}
	}

	if result == 0 {
		return ""
	}

	return string(result)
}