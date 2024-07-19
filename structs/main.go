package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	Name    string
	Age     int
	Contact Contact
}

// question 1 - Define a struct to represent a Person with fields like Name and Age , and write a function to print a Person .
func printStruct(p Person) {
	fmt.Printf("Name - %v, Age - %v\n", p.Name, p.Age)
}

// question 2 -Implement methods on a Person struct, suh as a method to check if the person is an adult.
func (p Person) isAdult() {
	if p.Age >= 18 {
		fmt.Println("Person is adult")
	} else {
		fmt.Println("Person is not an adult")
	}
}

// question 3 - Create a struct that contains nested structs and write a function to manipulate them.
func (e Employee) print() {
	fmt.Printf("Employee - Name - %v, Age - %v | Contact - Email - %v", e.Name, e.Age, e.Contact.Email)
}

// question 4 - Write a program to serialize and deserialize a struct to and from JSON.
func encodeDecode() {
	employee := Employee{Name: "Mobashir", Age: 19, Contact: Contact{
		Email: "a@b.com", Phone: "9874389120",
	}}

	// serialize JSON
	jsonBytes, _ := json.Marshal(employee)
	fmt.Println("serialized", string(jsonBytes))

	// deserialize from JSON
	var emp2 Employee
	json.Unmarshal(jsonBytes, &emp2)
	fmt.Println(emp2.Name, emp2.Age, emp2.Contact.Email, emp2.Contact.Phone)
}

func main() {
	// person := Person{Name: "Mobashir", Age: 15}
	// printStruct(person)
	// person.isAdult()
	// employee := Employee{Name: "Mobashir", Age: 19, Contact: Contact{
	//	Email: "a@b.com", Phone: "9874389120",
	//}}
	// employee.print()
	encodeDecode()
}
