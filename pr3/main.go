package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Условие:
Создай интерфейс Shape с методом Area() float64.
Реализуй две структуры: Circle и Rectangle.
Напиши функцию, которая принимает слайс []Shape и выводит площади всех фигур.
Добавь тип-ассерт для дополнительной логики: если фигура — круг, вывести “Круговая мощь!”.
*/

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func main() {
	circle := &Circle{
		Radius: 20,
	}

	rectangle := &Rectangle{
		Length: 10,
		Width:  15,
	}

	shapes := []Shape{circle, rectangle}
	GetShapes(shapes)

}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func GetShapes(shape []Shape) {
	for _, v := range shape {

		if c, ok := v.(*Circle); ok {
			fmt.Printf("Площадь фигуры %v = %f\n", reflect.TypeOf(c), c.Area())
			fmt.Printf("Круговая мощь!\n")
		}
		if r, ok := v.(*Rectangle); ok {
			fmt.Printf("Треугольная мощь!\n")
			fmt.Printf("Площадь фигуры %v = %f\n", reflect.TypeOf(r), r.Area())
		}
	}
}
