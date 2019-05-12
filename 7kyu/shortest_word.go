package codewars_probs

import "strings"

func FindShort(s string) int {
	shortestLen := 4294967296
	for _, word := range strings.Split(s, " ") {
		wordLen := len(string(word))
		if wordLen < shortestLen {
			shortestLen = wordLen
		}
	}
	return shortestLen
}
