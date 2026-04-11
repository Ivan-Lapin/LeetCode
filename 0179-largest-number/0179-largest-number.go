func largestNumber(nums []int) string {

	strs := make([]string, len(nums))

	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0] == "0" {
		return "0"
	}

	var result strings.Builder

	for _, num := range strs {
		result.WriteString(num)
	}

	return result.String()
}