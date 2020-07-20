package main

import (
	"fmt"
	"math"
)

var (
	rub            float32
	dol            float32
	dolRate        float32 = 70.67
	cathet1        float64
	cathet2        float64
	deposite       float64
	percent        float64
	termOfDeposite int
)

func main() {
	convert()
	rightTriangle()
	depositeCalc()
}

func convert() {
	fmt.Println("Введите сумму (в рублях) для конвертации в доллары:")
	fmt.Scan(&rub)
	dol = rub / dolRate
	fmt.Printf("Сумма в долларах: %.2f\n", dol)
}

func rightTriangle() {
	fmt.Println("Введите через пробел длины катетов прямоугольного треугольника:")
	fmt.Scanln(&cathet1, &cathet2)
	area := cathet1 * cathet2 / 2
	hypothenuse := math.Pow((cathet1*cathet1 + cathet2*cathet2), 0.5)
	perimetr := cathet1 + cathet2 + hypothenuse
	fmt.Printf("Площадь треугольника: %.2f;\nПериметр треугольника: %.2f;\nГипотенуза: %.2f.\n", area, perimetr, hypothenuse)
}

func depositeCalc() {
	fmt.Println("Введите сумму вклада:")
	fmt.Scanln(&deposite)
	fmt.Println("Введите срок вклада (полных лет):")
	fmt.Scanln(&termOfDeposite)
	fmt.Println("Введите годовой процент:")
	fmt.Scanln(&percent)
	// S = P*(1 + I/100)^n
	summ := deposite * math.Pow((1+percent*0.01), float64(termOfDeposite))
	fmt.Printf("Сумма вклада через %d лет, составит: %.2f\n", termOfDeposite, summ)
}
