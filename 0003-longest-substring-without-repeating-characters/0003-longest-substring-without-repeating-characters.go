func lengthOfLongestSubstring(s string) int {

	letterMap := make(map[byte]int)
	lenght := 0
	maxLen := 0

	l, r := 0, 0
	for r < len(s) {
		if _, exist := letterMap[s[r]]; exist {
			temp := l
			l = letterMap[s[r]] + 1
			for i := temp; i < letterMap[s[r]]; i++ {
				delete(letterMap, s[i])
			}
			l = letterMap[s[r]] + 1
			lenght = r - l + 1
			letterMap[s[r]] = r
		} else {
			letterMap[s[r]] = r
			lenght = r - l + 1
		}
		if lenght > maxLen {
			maxLen = lenght
		}
		r++
	}

	return maxLen
}