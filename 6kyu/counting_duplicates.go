package codewars_probs

import "strings"

func Duplicate_count(s1 string) int {
	appearanceMap := map[string]int{}
	counter := 0
	for _, letter := range s1 {
		strLetter := strings.ToLower(string(letter))
		if _, ok := appearanceMap[strLetter]; ok {
			appearanceMap[strLetter] += 1
		} else {
			appearanceMap[strLetter] = 1
		}
	}
	for _, val := range appearanceMap {
		if val > 1 {
			counter += 1
		}
	}
	return counter
}