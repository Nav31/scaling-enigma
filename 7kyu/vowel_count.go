package codewars_probs

func GetCount(str string) (count int) {
	count = 0
	for i := 0; i < len(str); i++ {
		letter := string(str[i])
		if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" {
			count += 1
		}
	}
	return count
}
