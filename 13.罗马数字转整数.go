func romanToInt(s string) int {
	s = strings.ReplaceAll(s, "IV", "a")
	s = strings.ReplaceAll(s, "IX", "b")
	s = strings.ReplaceAll(s, "XL", "c")
	s = strings.ReplaceAll(s, "XC", "d")
	s = strings.ReplaceAll(s, "CD", "e")
	s = strings.ReplaceAll(s, "CM", "f")
	
	result := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I': result += 1
		case 'V': result += 5
		case 'X': result += 10
		case 'L': result += 50
		case 'C': result += 100
		case 'D': result += 500
		case 'M': result += 1000
		case 'a': result += 4
		case 'b': result += 9
		case 'c': result += 40
		case 'd': result += 90
		case 'e': result += 400
		case 'f': result += 900
		}
	}
	return result
}
