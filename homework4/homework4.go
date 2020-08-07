package main

import (
	"fmt"
	"math"
)

type cardBoardBox struct {
	title               string
	length              float64
	height              float64
	depth               float64
	thickness           float64
	materialDensityKgM2 float64
}

type cup struct {
	title               string
	radius              float64
	height              float64
	thickness           float64
	materialDensityKgM2 float64
}

type shape interface {
	weight() float64
	surfaceArea() float64
}

func (c *cardBoardBox) weight() float64 {
	return c.surfaceArea() * c.materialDensityKgM2 * c.thickness
}

func (c *cardBoardBox) surfaceArea() float64 {
	return c.length*c.height*2 + c.length*c.depth*2 + c.height*c.depth*2
}

func (c *cup) weight() float64 {
	return c.surfaceArea() * c.materialDensityKgM2 * c.thickness
}

func (c *cup) surfaceArea() float64 {
	return c.radius*c.radius*math.Pi + 2*math.Pi*c.radius*c.height
}

func weightForShape(s shape) float64 {
	return s.weight()
}

func main() {
	korobka := &cardBoardBox{
		title:               "Картонная коробка",
		length:              0.5,
		height:              0.485,
		depth:               0.395,
		thickness:           0.06,
		materialDensityKgM2: 0.592,
	}
	teaCup := &cup{
		title:               "Чайная кружка",
		radius:              0.03,
		height:              0.085,
		thickness:           0.027,
		materialDensityKgM2: 1430,
	}
	fmt.Printf("%v, вес составляет: %.3f\n", korobka.title, weightForShape(korobka))
	fmt.Printf("%v, вес составляет: %.3f\n", teaCup.title, weightForShape(teaCup))
}
