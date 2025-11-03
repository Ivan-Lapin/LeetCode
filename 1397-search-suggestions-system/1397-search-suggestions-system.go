func suggestedProducts(products []string, searchWord string) [][]string {
    
    partWords := make(map[string][]string)
    result := [][]string{}
    for _, product := range products {
        for i := 1; i <= len(product); i++{
            part := product[:i]
            if _, exist := partWords[part]; exist {
                partWords[part] = SortInsert(partWords[part], product)
            } else {
                partWords[part] = []string{product}
            }
        }
    }

    for i := 1; i <= len(searchWord); i++{
        part := searchWord[:i]
        if len(partWords[part]) > 3 {
            result = append(result, partWords[part][:3])
        } else {
            result = append(result, partWords[part])
        }
    }

    return result

}

func SortInsert(words []string, str string) []string {
    new_words := make([]string, len(words)+1)
    for i := 0; ; i++{

        if i == len(words) {
            new_words[i] = str
            break
        }

        if words[i] < str {
            new_words[i] = words[i]
        } else {
            new_words[i] = str
            for j := i; j < len(words); j++{
                new_words[j+1] = words[j]
            }
            break
        }
    }
    return new_words
}