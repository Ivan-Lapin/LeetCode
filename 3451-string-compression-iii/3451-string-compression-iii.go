func compressedString(word string) string {

	var comp strings.Builder
	comp.Grow(len(word) * 2)

	count := byte(1)
	symbol := word[0]

	for i := 1; i < len(word); i++ {

		if word[i] != symbol || count >= 9 {

			comp.WriteByte('0' + count)
			comp.WriteByte(symbol)

			symbol = word[i]
			count = byte(1)

		} else {
			count++
		}
	}
	comp.WriteByte('0' + count)
	comp.WriteByte(symbol)

	return comp.String()
}