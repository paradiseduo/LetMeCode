func isPalindrome(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	var numArr []int
	for {
		if x == 0 {
			break
		}
		numArr = append(numArr, x%10)
		x = x / 10
	}
	flag := true
	if len(numArr)%2 == 1 {
		for i := 0; i < len(numArr); i++ {
			if i == (len(numArr)+1)/2 {
				return flag
			}
			if numArr[i] != numArr[len(numArr)-(i+1)] {
				return false
			}
		}
	} else {
		for i := 0; i < len(numArr); i++ {
			if i == len(numArr)/2 {
				return flag
			}
			if numArr[i] != numArr[len(numArr)-(i+1)] {
				return false
			}
		}
	}
	return flag
}
