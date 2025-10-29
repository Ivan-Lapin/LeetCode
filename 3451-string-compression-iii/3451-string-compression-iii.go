func compressedString(word string) string {

	var comp strings.Builder
	count := 1
	symbol := word[0]

	for i := 1; i < len(word); i++ {

		if word[i] != symbol || count >= 9 {

			comp.WriteString(fmt.Sprint(count, string(symbol)))
			symbol = word[i]
			count = 1

		} else {
			count++
		}
	}
	comp.WriteString(fmt.Sprint(count, string(symbol)))

	return comp.String()
}