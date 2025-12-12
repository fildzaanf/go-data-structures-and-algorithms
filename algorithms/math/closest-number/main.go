package main

import (
	"fmt"
	"math"
)

func ClosestNumber(target, divisor int) int {

	lowerMultiple := (target / divisor) * divisor

	upperMultiple := 0

	if target%divisor == 0 {
		upperMultiple = lowerMultiple
	} else if target > 0 {
		upperMultiple = lowerMultiple + divisor
	} else {
		upperMultiple = lowerMultiple - divisor
	}

	distanceToLower := int(math.Abs(float64(target - lowerMultiple)))
	distanceToUpper := int(math.Abs(float64(target - upperMultiple)))

	if distanceToLower < distanceToUpper {
		return lowerMultiple
	} else if distanceToLower > distanceToUpper {
		return upperMultiple
	}

	if int(math.Abs(float64(lowerMultiple))) > int(math.Abs(float64(upperMultiple))) {
		return lowerMultiple
	}

	return upperMultiple
}

func main() {
	target := -15
	divisor := 6

	result := ClosestNumber(target, divisor)
	fmt.Print(result)
}
