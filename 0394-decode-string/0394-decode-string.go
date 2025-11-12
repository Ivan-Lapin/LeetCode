func decodeString(s string) string {
	var result strings.Builder

	var def func(number int, index int) (string, int)
	var define_number func(index int) (int, int)

	def = func(number int, index int) (string, int) {

		var output_sub strings.Builder

		for i := index; i < len(s); i++ {

			if s[i] >= 48 && s[i] <= 57 {

				num, next := define_number(i)

				var next_sub string
				next_sub, next = def(num, next)
                
				output_sub.WriteString(next_sub)
				i = next

				continue
			}

			if s[i] == ']' {
				index = i
				break
			}

			output_sub.WriteByte(s[i])
		}

		temp := output_sub.String()
		for i := 1; i < number; i++ {
			output_sub.WriteString(temp)
		}

		return output_sub.String(), index
	}

	define_number = func(index int) (int, int) {

		var number strings.Builder

		for s[index] >= 48 && s[index] <= 57 {
			number.WriteByte(s[index])
			index++
		}

		num, _ := strconv.Atoi(number.String())
		index++

		return num, index
	}

	for i := 0; i < len(s); i++ {

		if s[i] >= 49 && s[i] <= 57 {

			number, next_index := define_number(i)

			var sub string
			sub, next_index = def(number, next_index)

			result.WriteString(sub)
			i = next_index

			continue
		}
        
		if s[i] != '[' && s[i] != ']' {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}