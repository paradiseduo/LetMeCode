func checkXMatrix(grid [][]int) bool {
	count := len(grid)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i == j || i+j+1 == count {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}

	return true
}
