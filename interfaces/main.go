package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
}

type Rectangle struct {
	width, heigth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.heigth
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// func calculateArea(s Shape) float64 {
// 	return s.Area()
// }

func main() {
	rect := Rectangle{width: 5, heigth: 4}
	circle := Circle{radius: 2}
	fmt.Println(rect, circle)
	// fmt.Println("Rectangle area is:", calculateArea(rect)) COMPILER ERROR
	// fmt.Println("Rectangle area is:", calculateArea(circle)) COMPILER ERROR
}
