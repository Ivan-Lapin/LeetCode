func countWords(words1 []string, words2 []string) int {

    freqWord1Map := make(map[string]int)
    freqWord2Map := make(map[string]int)

    count := 0

    for _, word := range words1 {
        freqWord1Map[word]++
    }

    for _, word := range words2 {
        freqWord2Map[word]++
    }

    for key, val := range freqWord1Map {
        if freqWord2Map[key] == 1 && val == 1 {
            count++
        }
    }

    return count
}