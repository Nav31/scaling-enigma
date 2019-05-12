package codewars_probs

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	const seperator = "-"
	var built_strs []string
	splitStr := strings.Split(s, "")
	for i := 0; i < len(splitStr); i++ {
		letter := splitStr[i]
		fmt.Println(letter)
		built_strs = append(built_strs, strings.Repeat(letter, i+1))
		fmt.Println(built_strs)
	}
	return ""
}