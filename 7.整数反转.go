func reverse(x int) int {
	n := 0
	for true {
		n = (x % 10) + n*10
		x = x / 10
		if x == 0 {
			break
		}
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}
	return n
}
