package main

import (
	"errors"
	"fmt"
	"math"
)

// question 1 : Classes Equivalent - Define a struct Employee with fields name , age , and a method displayDetails() to print employee details.
type Employee struct {
	Name string
	Age  int
	Post string
}

func (e Employee) Greet() {
	fmt.Printf("Name - %s | Age - %d | Post - %s", e.Name, e.Age, e.Post)
}

// question 2 : Inheritance Equivalent - Create a struct Vehicle with a method start() and embed it in structs Car and Bike , overriding start() in both.
type Vehicle struct{}

func (v Vehicle) start() {
	fmt.Println("Vehicle is starting")
}

// Car struct embedding Vehicle
type Car struct {
	Vehicle
}

func (c Car) start() {
	fmt.Println("Car is starting")
}

// Bike struct embedding Vehicle
type Bike struct {
	Vehicle
}

func (b Bike) start() {
	fmt.Println("Bike is starting")
}

// question 3 - Polymorphism - Define an interface Shape with a method Area() . Implement Area() for structs Circle and Rectangle .
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

// question - 4 - Encapsulation - Create a struct BankAccount with unexported fields balance and exported methods Deposit() and Withdraw()
type BankAccount struct {
	balance float64
}

// Deposit method adds the amount to the balance
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	b.balance += amount
	return nil
}

// Withdraw method subtracts the amount from the balance
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > b.balance {
		return errors.New("insufficient funds")
	}
	b.balance -= amount
	return nil
}

// Balance method returns the current balance
func (b *BankAccount) Balance() float64 {
	return b.balance
}

// question - 5 - Abstract Class -Define an interface Animal with a default method Eat() . Implement Eat() for structs Dog and Cat .
// Animal interface with Eat method
type Animal interface {
	Eat()
}

type Dog struct {
	Name string
}

func (d Dog) Eat() {
	fmt.Println(d.Name, "is eating dog food.")
}

type Cat struct {
	Name string
}

func (c Cat) Eat() {
	fmt.Println(c.Name, "is eating cat food.")
}

func main() {
	fmt.Println("Hello from oop")

	//employee := Employee{
	//	Name: "Mobashir Farhan",
	//	Age:  20,
	//	Post: "CEO",
	// }

	// employee.Greet()

	// Create instances of Car and Bike
	myCar := Car{}
	myBike := Bike{}

	// Call the start method on Car and Bike
	myCar.start()  // Output: Car is starting
	myBike.start() // Output: Bike is starting

	// You can still access the Vehicle's start method directly if needed
	myCar.Vehicle.start()  // Output: Vehicle is starting
	myBike.Vehicle.start() // Output: Vehicle is starting

	// Create instances of Dog and Cat
	myDog := Dog{Name: "Rex"}
	myCat := Cat{Name: "Whiskers"}

	// Create a slice of Animal interface
	animals := []Animal{myDog, myCat}

	// Call the Eat method for each Animal
	for _, animal := range animals {
		animal.Eat()
	}
}
