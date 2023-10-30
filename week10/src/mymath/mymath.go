package mymath

func MyAbs(number int) int { // 절대 값 함수
	if number < 0 { // 음수면
		return number * -1
	}
	return number
}
