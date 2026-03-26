func topKFrequent(nums []int, k int) []int {

    mapa := make(map[int]int)
    frequents := []int{}
    result := []int{}

    for _, num := range nums {
        mapa[num]++
    }

    for _, val := range mapa{
        frequents = append(frequents, val)
    }

    sort.Slice(frequents, func(i, j int) bool {
			return frequents[j] < frequents[i]
	})

    for i := range k{
        freq := frequents[i]

        for key, val := range mapa {
            if val == freq {
                result = append(result, key)
                delete(mapa, key)
                break
            }
        }
    }

    return result
}