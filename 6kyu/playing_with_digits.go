package codewars_probs

import (
    "math"
    "strconv"
)

// https://www.codewars.com/kata/playing-with-digits/train/go

func DigPow(n, p int) int {
    strNum := strconv.Itoa(n)
    intermediateNum := 0.0
    for idx, val := range strNum {
        digit, _ := strconv.Atoi(string(val))
        a := float64(digit)
        b := float64(p + idx)
        intermediateNum += math.Pow(a, b)
    }
    answer := intermediateNum / float64(n)
    if math.Mod(answer, 1) == 0 {
        return int(answer)
    } else {
        return -1
    }
}
