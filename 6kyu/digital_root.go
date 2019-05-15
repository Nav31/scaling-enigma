package codewars_probs

import (
	"strconv"
)

// https://www.codewars.com/kata/541c8630095125aba6000c00

func innerFunc(stringDigit string) int {
	runningSum := 0
	for _, numb := range stringDigit {
		intNumb, _ := strconv.Atoi(string(numb))
		runningSum += intNumb
	}
	return runningSum
}

func DigitalRoot(n int) int {
	for n > 9 {
		n = innerFunc(strconv.Itoa(n))
		print(n)
	}
	return n
}
