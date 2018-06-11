package util

func BoolToInt(flag bool) int {
	if flag {
		return 1
	}

	return 0
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
