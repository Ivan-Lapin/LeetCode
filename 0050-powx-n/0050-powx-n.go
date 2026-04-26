func myPow(x float64, n int) float64 {

	flag := true
	if n < 0 {
		flag = false
		n = -n
	}

	result := 1.0

	for n > 0 {
		if n%2 == 0 {
			n /= 2
			x *= x
		} else {
			n--
			result *= x
		}
	}

	if flag {
		return result
	}
	return 1 / result
}