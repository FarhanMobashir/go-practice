package main

import "fmt"

// question 1 - Write a function to swap the values of two integers using pointers.
func swap(int1 *int, int2 *int) {
	temp := *int1
	*int1 = *int2
	*int2 = temp
}

// question 2 - Create a function that modifies an array in place using pointers.
func modifyArrayBy1(arr *[]int) {
	for i := range *arr {
		(*arr)[i]++
	}
}

// question 3 - Write a function that checks if a pointer is nil and handles it appropriately.
func checkNilPointer(ptr *int) {
	if ptr == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}
}

type Person struct {
	Name string
	Age  int
}

// question 4 - Write a function that modifies a struct by passing it by reference.
func modifyStruct(P *Person) {
	P.Name = "John Doe"
	P.Age = 20
}

func main() {
	fmt.Print("Hello from pointers")
	a := 5
	b := 10
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	array := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", array)

	modifyArrayBy1(&array)
	fmt.Println("Modified array:", array)

	var p *int
	var q int = 10
	r := &q

	checkNilPointer(p) // Should print "Pointer is nil"
	checkNilPointer(r) // Should print "Pointer is not nil"

	person := Person{Name: "Jane Smith", Age: 25}
	fmt.Println("Before modification:", person)

	// Call the function to modify the struct
	modifyStruct(&person)
	fmt.Println("After modification:", person)
}
