package main

import "fmt"

var (
	number int
	fibo   [100]uint
)

func main() {
	evenNumber()
	partibilityThree()
	hundredElemsOfFibonacci()
}

func evenNumber() {
	fmt.Println("Введите число:")
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Printf("Число %d четное\n", number)
	} else {
		fmt.Printf("Число %d нечетное\n", number)
	}
}

func partibilityThree() {
	if number%3 == 0 {
		fmt.Printf("Число %d делится на 3 без остатка\n", number)
	} else {
		fmt.Printf("Число %d не делится на 3 без остатка\n", number)
	}
}

func hundredElemsOfFibonacci() {
	fmt.Println("Первые 100 чисел Фибоначчи:")
	fibo[0] = 0
	fibo[1] = 1
	fmt.Println(fibo[0])
	fmt.Println(fibo[1])
	for i := 2; i < 100; i++ {
		fibo[i] = fibo[i-2] + fibo[i-1]
		fmt.Println(fibo[i])
	}
}
