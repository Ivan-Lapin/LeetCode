func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    w1 := strings.Join(word1, "")
    w2 := strings.Join(word2, "")

    return w1 == w2
}