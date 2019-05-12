package codewars_probs

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
  highestSeen := -4294967296
  lowestSeen := 4294967296
  retStr := ""
  newStr := strings.Split(in, " ")
  for _, strNumb := range newStr {
  	potentialNumb := string(strNumb)
  	numb, _ := strconv.Atoi(string(potentialNumb))
  	if numb > highestSeen {
  		highestSeen = numb
	}
  	if numb < lowestSeen {
		lowestSeen = numb
	}
  }
  return retStr + strconv.Itoa(highestSeen) + " " + strconv.Itoa(lowestSeen)
}