package main

import (
	"fmt"
	"math"
)

func DistanceBetweemTwoPoints(x1, x2, y1, y2 int) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)

	d := math.Sqrt((dx * dx) + (dy * dy))

	return d
}

func main() {
	x1, y1 := 3, 4
	x2, y2 := 7, 7
	fmt.Printf("Distance: %.5f\n", DistanceBetweemTwoPoints(x1, x2, y1, y2))

	x1, y1 = 3, 4
	x2, y2 = 4, 3
	fmt.Printf("Distance: %.5f\n", DistanceBetweemTwoPoints(x1, x2, y1, y2))
}
