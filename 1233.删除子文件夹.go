func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	result := []string{folder[0]}
	for i := 1; i < len(folder); i++ {
		cur := result[len(result)-1]
		if !strings.HasPrefix(folder[i], cur) || folder[i][len(cur)] != '/' {
			result = append(result, folder[i])
		}
	}
	return result
}
