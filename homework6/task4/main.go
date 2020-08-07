package main

import (
	"fmt"
	"strconv"

	"github.com/rockstar90/go_course/Go-Course/homework6/task4/quadratic"
)

var input string
var num []float64
var nameCoefficient = []string{"a", "b", "c"}

func main() {
	fmt.Println("Программа предназначена для решения квадратных уравнений вида: ax^2+bx+c=0")
	for i := 0; i < 3; i++ {
		fmt.Printf("Введите значение коэффициента %v:\n", nameCoefficient[i])
		fmt.Scanln(&input)
		value, _ := strconv.ParseFloat(input, 64)
		num = append(num, value)
	}
	x1, x2 := quadratic.QuadraticRoots(num)
	fmt.Println("Корни:", x1, x2)
}
