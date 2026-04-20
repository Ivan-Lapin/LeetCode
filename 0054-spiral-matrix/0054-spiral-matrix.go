func spiralOrder(matrix [][]int) []int {

	size := len(matrix) * len(matrix[0])
	result := make([]int, 0, size)
	r_down, r_up := len(matrix)-1, 1
	c_left, c_right := 0, len(matrix[0])-1

	for i, j := 0, -1; ; {
		for j < c_right {
			j++
			result = append(result, matrix[i][j])
		}
		c_right = j - 1

		if len(result) == size {
			break
		}

		for i < r_down {
			i++
			result = append(result, matrix[i][j])
		}
		r_down = i - 1

		if len(result) == size {
			break
		}

		for j > c_left {
			j--
			result = append(result, matrix[i][j])
		}
		c_left = j + 1

		if len(result) == size {
			break
		}

		for i > r_up {
			i--
			result = append(result, matrix[i][j])
		}
		r_up = i + 1

		if len(result) == size {
			break
		}
	}

	return result
}