func dieSimulator(n int, rollMax []int) int {
	mod := 1000000007
	d := [5001][6]int{}
	sum := [5001]int{}
	sum[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < 6; j++ {
			pos := 0
			if i-rollMax[j]-1 > 0 {
				pos = i - rollMax[j] - 1
			}
			sub := ((sum[pos]-d[pos][j])%mod + mod) % mod
			d[i][j] = ((sum[i-1]-sub)%mod + mod) % mod
			if i <= rollMax[j] {
				d[i][j] = (d[i][j] + 1) % mod
			}
			sum[i] = (sum[i] + d[i][j]) % mod
		}
	}
	return sum[n]
}
