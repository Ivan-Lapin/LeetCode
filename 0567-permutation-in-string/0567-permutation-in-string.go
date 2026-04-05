func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	m := make(map[rune]int)
	accordance := 0

	for _, el := range s1 {
		m[el]++
	}

	for i := range len(s1) {
		el := rune(s2[i])
		if _, exist := m[el]; exist {
			m[el]--
            if m[el] < 0 {
                accordance--
            } else {
                accordance++
            }
		}
	}

	if accordance == len(s1) {
		return true
	}
	l := 1
	r := l + len(s1) - 1
	for r < len(s2) {
		
		past := rune(s2[l-1])
		if _, exist := m[past]; exist {
			m[past]++
            if m[past] <= 0 {
                accordance++
            } else {
                accordance--
            } 
		}

		new := rune(s2[r])
		if _, exist := m[new]; exist {
			m[new]--
            if m[new] < 0 {
                accordance--
            } else {
                accordance++
            }
		}


		if accordance == len(s1) {
			return true
		}

		l++
		r++

	}

	return false
}