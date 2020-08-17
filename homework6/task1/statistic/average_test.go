package statistic

import (
	"fmt"
	"testing"
)

type testValue struct {
	inputVal  []float64
	outputVal float64
	sumInputs float64
}

var tests = []testValue{
	{[]float64{1, 1, 6, 1, 1}, 2, 10},
	{[]float64{1, 2, 1, 2, 1, 2}, 1.5, 9},
	{[]float64{1, 2, 3, 4, 5}, 3, 15},
}

func TestAverage(t *testing.T) {
	for i, pair := range tests {
		value := pair.outputVal
		if value != Average(pair.inputVal) {
			fmt.Printf("Где-то в коде косяк! Не прошел тест №%v", i+1)
		}
	}
}

func TestSumSlice(t *testing.T) {
	for i, pair := range tests {
		value := pair.sumInputs
		if value != SumSlice(pair.inputVal) {
			fmt.Printf("Где-то в коде косяк! Не прошел тест №%v\n", i+1)
		}
	}
}
