package codewars_probs

import (
	"sort"
	"strings"
)

//https://www.codewars.com/kata/two-to-one/train/go

func TwoToOne(s1, s2 string) string {
	combinedStr := s1 + s2
	stringMap := map[string]string{}
	stringSlice := []string{}
	for _, charCode := range combinedStr {
		letter := string(charCode)
		if _, ok := stringMap[letter]; !ok {
			stringMap[letter] = letter
		}
	}
	for key, _ := range stringMap {
		stringSlice = append(stringSlice, key)
	}
	sort.Strings(stringSlice)
	return strings.Join(stringSlice, "")
}