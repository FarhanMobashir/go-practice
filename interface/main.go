package main

import (
	"fmt"
	"math"
)

// question 1- Define an interface for a Shape with methods Area() and Perimeter() , and implement it for Circle and Rectangle .
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Width) * float64(r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// question 2 - Write a function that takes an empty interface and performs type assertion to determine the underlying type.
func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case float64:
		fmt.Println("Float", v)
	default:
		fmt.Println("Unknown type")
	}
}

// question 3 - Implement a function that takes an empty interface and handles different data types.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	fmt.Println("Hello from interface")

	// circle := Circle{Radius: 5}
	// rectangle := Rectangle{Width: 4, Height: 3}

	// fmt.Printf("Circle Area: %f\n", circle.Area())
	// fmt.Printf("Circle Perimeter: %f\n", circle.Perimeter())
	// fmt.Printf("Rectangle Area: %f\n", rectangle.Area())
	// fmt.Printf("Rectangle Perimeter: %f\n", rectangle.Perimeter())
	// printType(10)
	// printType(3.14)
	// printType("Hello")
	// printType(true)

	describe(10)             // Output: (10, int)
	describe(3.14)           // Output: (3.14, float64)
	describe("Hello")        // Output: (Hello, string)
	describe(true)           // Output: (true, bool)
	describe([]int{1, 2, 3}) // Output: ([1 2 3], []int)
}
