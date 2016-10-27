package main

import (
	"fmt"
	"math"
)

func relativeDifference(a, b float64) float64 {

	return math.Abs(a-b) / a

}

func Sqrt(x float64) float64 {

	var z float64 = 1.00

	for i, previousZ := 1, 0.00; relativeDifference(z, previousZ) >= 0.001 && i < 1000; i++ {

		previousZ = z
		z = (z - ((z*z - x) / (2 * z)))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(0.01))
}
