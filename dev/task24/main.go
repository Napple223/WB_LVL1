package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния
//между двумя точками, которые представлены
//в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.

// Тип данных точка.
type Point struct {
	x float64
	y float64
}

// Функция-конструктор, возвращающая точку.
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Функция для нахождения расстояние от одной точки до другой.
func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

func main() {
	p1 := NewPoint(3, -6)
	p2 := NewPoint(-5, 7)
	fmt.Println(p1.Distance(p2) == p2.Distance(p1))
}
