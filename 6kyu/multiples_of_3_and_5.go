package codewars_probs

import "math"

func Multiple3And5(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		if math.Mod(float64(i), 3) == 0 && math.Mod(float64(i), 5) == 0 {
			sum += i
		} else if math.Mod(float64(i), 3) == 0 {
			sum += i
		} else if math.Mod(float64(i), 5) == 0 {
			sum += i
		}
	}
	return sum
}

// https://www.codewars.com/kata/multiples-of-3-or-5/train/go

