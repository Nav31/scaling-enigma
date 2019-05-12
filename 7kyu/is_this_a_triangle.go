package codewars_probs

func IsTriangle(a, b, c int) bool {
	if (a + b) <= c {
		return false
	} else if (a + c) <= b {
		return false
	} else if (b + c) <= a {
		return false
	} else {
		return true
	}
}
