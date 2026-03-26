func groupAnagrams(strs []string) [][]string {

	mapa := make(map[string][]string)
	for i := range strs {
		runes := []rune(strs[i])

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		word := string(runes)
		mapa[word] = append(mapa[word], strs[i])
	}

	result := [][]string{}
	for _, val := range mapa {
		result = append(result, val)
	}

	return result
}