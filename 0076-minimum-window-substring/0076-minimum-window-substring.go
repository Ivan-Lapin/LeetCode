func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	} else if t == s {
		return s
	}

	m := make(map[rune]int)

	for i := range len(t) {
		el := rune(t[i])
		m[el]++
	}

    var result string
    minLen := len(s)
	accordance := 0
	miss := make([]int, 0, len(s))

	start := 0
	for start < len(s) {
		if _, exist := m[rune(s[start])]; exist {
			break
		} else {
			start++
		}
	}

	l := start
	r := l
    var el rune

	for r < len(s) {

		el = rune(s[r])
		if _, exist := m[el]; exist {
			m[el]--
			if m[el] >= 0 {
				accordance++
			}
			if r != l {
				miss = append(miss, r)
			}
		}

		for accordance == len(t) {

            if len(t) == 1 {
				return t
			}

			if len(s[l:r]) < minLen {
				minLen = len(s[l : r+1])
				result = s[l : r+1]
			}

			if len(miss) != 0 {
				m[rune(s[l])]++
				if m[rune(s[l])] > 0 {
					accordance--
				}
				l = miss[0]
				miss = miss[1:]
			}
		}

		r++
	}

	return result
}