func finalValueAfterOperations(operations []string) int {
	result := 0
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "X++", "++X":
			result += 1
		case "X--", "--X":
			result -= 1
		}
	}
	return result
}
