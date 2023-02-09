type AuthenticationManager struct {
	timeLine int
	code     map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	m := AuthenticationManager{
		timeLine: timeToLive,
		code:     map[string]int{},
	}
	return m
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.code[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	v, e := this.code[tokenId]
	if e && v+this.timeLine > currentTime {
		this.code[tokenId] = currentTime
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	result := 0
	for _, v := range this.code {
		if v+this.timeLine > currentTime {
			result += 1
		}
	}
	return result
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
