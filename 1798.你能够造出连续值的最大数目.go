func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	x := 1
	for i := 0; i < len(coins); i++ {
		if coins[i] <= x {
			x += coins[i]
		} else {
			break
		}
	}
	return x
}
