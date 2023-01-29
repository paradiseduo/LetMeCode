func lengthOfLongestSubstring(s string) int {
	lens := len(s)
	window := make(map[byte]int, lens)
	left, right, res := 0, 0, 0
	for right < lens {
		b := s[right]
		window[b] += 1
		right += 1
		for window[b] > 1 {
			c := s[left]
			window[c] -= 1
			left += 1
		}
		if right - left > res {
			res = right - left
		}
	}
	return res
}
