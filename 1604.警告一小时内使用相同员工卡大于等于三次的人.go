func alertNames(keyName []string, keyTime []string) []string {
	result := []string{}
	users := map[string][]int{}
	for i := 0; i < len(keyName); i++ {
		ss := strings.Split(keyTime[i], ":")
		hour, _ := strconv.Atoi(ss[0])
		min, _ := strconv.Atoi(ss[1])
		time := hour*60 + min
		if users[keyName[i]] != nil {
			users[keyName[i]] = append(users[keyName[i]], time)
		} else {
			users[keyName[i]] = []int{time}
		}
	}
	for key, value := range users {
		if len(value) >= 3 {
			sort.Ints(value)
			for i := 2; i < len(value); i++ {
				if value[i]-value[i-2] <= 60 {
					result = append(result, key)
					break
				}
			}
		}
	}
	sort.Strings(result)
	return result
}
