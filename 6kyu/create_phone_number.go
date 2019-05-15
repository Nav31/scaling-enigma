package codewars_probs

import "strconv"

// https://www.codewars.com/kata/create-phone-number/train/go

func CreatePhoneNumber(numbers [10]uint) string {
	retStr := ""
	for idx, val := range numbers {
		if idx == 0 {
			retStr += "("
		} else if idx == 3 {
			retStr += ") "
		} else if idx == 6 {
			retStr += "-"
		}
		retStr += strconv.Itoa(int(val))
	}
	return retStr
}
