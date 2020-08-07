package quadratic

import (
	"fmt"
	"math"
)

func QuadraticRoots(coeffs []float64) (x1, x2 float64) {
	discrim := math.Pow(coeffs[1], 2) - 4*coeffs[0]*coeffs[2]
	if discrim < 0 {
		fmt.Println("Корней на множестве действительных чисел нет")
		return x1, x2
	} else if discrim == 0 {
		x1 := -coeffs[1] / (2 * coeffs[0])
		x2 := x1
		fmt.Printf("Существует единственный корень: %.1f\n", x1)
		return x1, x2
	} else {
		x1 = (-coeffs[1] + math.Sqrt((math.Pow(coeffs[1], 2) - 4*coeffs[0]*coeffs[2]))) / (2 * coeffs[0])
		x2 = (-coeffs[1] - math.Sqrt((math.Pow(coeffs[1], 2) - 4*coeffs[0]*coeffs[2]))) / (2 * coeffs[0])
		fmt.Printf("Корней два:\nX1: %.1f\nX2: %.1f\n", x1, x2)
		return x1, x2
	}
}
