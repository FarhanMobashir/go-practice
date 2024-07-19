package main

import "fmt"

// question 1 - Write a function that returns multiple values, e.g., the quotient and remainder of two integers.
func quotientAndRemainder(int1 int, int2 int) (int, int) {
	quotient := int1 / int2
	remainder := int1 % int2

	return quotient, remainder
}

// question 2 - Implement a function that takes another function as an argument and applies it to a value.
func apply(f func(int) int, x int) int {
	return f(x)
}

func double(n int) int {
	return n * 2
}

// question 3 - Write a recursive function to compute the factorial of a number.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

// question 4 - Implement a function that applies a given function to all elements in an array.
func mapFunction(array []int, f func(int) int) []int {
	result := make([]int, len(array))

	for i, v := range array {
		result[i] = f(v)
	}

	return result
}

func filterFunction(array []int, f func(int) bool) []int {
	result := []int{}

	for _, v := range array {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {

	// q, r := quotientAndRemainder(16, 3)
	// fmt.Printf("Q - %d, R - %d", q, r)
	// num := 5
	// result1 := apply(double, num)
	// fmt.Printf("Applying double to %d gives %d\n", num, result1)
	//fmt.Printf("Factorial of %d is %d\n", 5, factorial(5))
	// Example usage with an array of integers
	intArray := []int{1, 2, 3, 4, 5}

	// Filter function to keep only even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	evenNumbers := filterFunction(intArray, isEven)
	fmt.Println("Even numbers:", evenNumbers)

	// Example usage with an array of integers
	double := func(n int) int {
		return n * 2
	}
	doubledIntArray := mapFunction(intArray, double)
	fmt.Println("Doubled integers:", doubledIntArray)

}
