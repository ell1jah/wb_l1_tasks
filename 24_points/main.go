package main

import (
	"fmt"
	"math"
)

type Point struct {
	// пишем с маленькой для сокрытия
	x float64
	y float64
}

// конструктор
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(to *Point) float64 {
	// по т. Пифагора находим расстояние
	return math.Sqrt((p.x-to.x)*(p.x-to.x) + (p.y-to.y)*(p.y-to.y))
}

func main() {
	// создаем точки
	a := NewPoint(1.5, 2.4)
	b := NewPoint(4.5, 6.4)

	fmt.Println("distance: ", a.Distance(b))
}
