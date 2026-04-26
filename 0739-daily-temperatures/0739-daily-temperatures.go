func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
    deque := []int{}

	for i := 0; i < len(temperatures); i++{
		
        for len(deque) > 0 && temperatures[i] > temperatures[deque[len(deque)-1]] {
            res[deque[len(deque)-1]] = i - deque[len(deque)-1]
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
	}

	return res
}