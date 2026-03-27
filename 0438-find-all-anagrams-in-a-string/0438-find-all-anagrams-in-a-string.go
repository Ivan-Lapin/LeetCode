func findAnagrams(s string, p string) []int {

    if len(p) > len(s) {
		return []int{}
	}

	smap := make(map[byte]int)
	result := []int{}

	for i := range len(p) {
		smap[p[i]]++
	}

	size_window := len(p)
	window := s[0:size_window]
	count := 0

	for i := 0; i < len(window); i++ {
		if val, exist := smap[window[i]]; exist {
			smap[window[i]]--
			if val > 0 {
				count++
			}
		}
	}

	if count == size_window {
		result = append(result, 0)
	}

	for i := 1; i < len(s)-(len(p)-1); i++ {

		if val, exist := smap[s[i-1]]; exist {
			if val >= 0 {
				count--
			}
			smap[s[i-1]]++
		}

		window := s[i : i+size_window]

		if val, exist := smap[window[size_window-1]]; exist {
			if val > 0 {
				count++
                if count > size_window {
                    count = size_window
                }
			}
			smap[window[size_window-1]]--
		}

		if count == size_window {
			result = append(result, i)
		}

	}

	return result
}