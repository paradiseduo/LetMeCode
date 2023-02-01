func decodeMessage(key string, message string) string {
	cur := byte('a')
	rule := map[byte]byte{}
	rule[' '] = ' '
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && rule[key[i]] == 0 {
			rule[key[i]] = cur
			cur += 1
		}
	}

	m := []byte(message)
	for i := 0; i < len(m); i++ {
		m[i] = rule[m[i]]
	}
	return string(m)
}
