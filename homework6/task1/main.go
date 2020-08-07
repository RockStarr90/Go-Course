package main

import (
	"fmt"

	"github.com/rockstar90/go_course/Go-Course/homework6/homework6p1/statistic"
)

var input float64 = 1
var inputSlice []float64

func main() {
	for input != 0 {
		fmt.Scanln(&input)
		if input != 0 {
			inputSlice = append(inputSlice, input)
		}
	}
	fmt.Println(statistic.Average(inputSlice))
}
