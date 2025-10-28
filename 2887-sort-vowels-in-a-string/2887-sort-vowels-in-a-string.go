func sortVowels(s string) string {
	vowelsMap := map[byte]bool{
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	} 

	vowels := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if vowelsMap[s[i]] {
			vowels = append(vowels, s[i])
		}
	}

	sort.Slice(vowels, func(i, j int) bool { return vowels[i] < vowels[j] }) 

	var b strings.Builder
	b.Grow(len(s)) 
	for i, j := 0, 0; i < len(s); i++ {
		if vowelsMap[s[i]] {
			b.WriteByte(vowels[j])
			j++ 
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
