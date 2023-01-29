func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		pIndex, value := m[target-val]
		if value {
			return []int{pIndex, index}
		}
		m[val] = index
	}
	return []int{}
}
