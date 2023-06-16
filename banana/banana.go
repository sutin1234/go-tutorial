package banana

import (
	"math"
)

func Solution(S string) float64 {
	letCount := make(map[string]int)
	bCount := map[string]int{"A": 1, "B": 3, "N": 2}
	maxLetter := math.MaxFloat64

	for _, char := range S {
		letter := char
		letCount[string(letter)] = letCount[string(letter)] + 1
	}
	for char := range bCount {
		lc := float64(letCount[char])
		bc := float64(bCount[char])
		occur := lc / bc
		maxLetter = math.Floor(occur)
	}

	return maxLetter
}
