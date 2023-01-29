func countAsterisks(s string) int {
	arr := strings.Split(s, "|")
	count := 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			sub := arr[i]
			for j := 0; j < len(sub); j++ {
				if sub[j] == '*' {
					count += 1
				}
			}
		}
	}
	return count
}
