package main

import "fmt"

// question 1 - Write a function that calculates the sum of the first N natural numbers.
func sumOfN(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	fmt.Printf("The sum is %d", sum)

}

// question 2 - Implement a function to compute the factorial of a given number.
func factorial(num int) {
	result := 1

	for i := 1; i < num; i++ {
		result *= i
	}

	fmt.Printf("The factorial is %d", result)
}

// question 3 - Write a program to print the first N Fibonacci numbers.
func fib(n int) {
	if n == 0 {
		return
	}

	var a, b int = 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()

}

// question 4 - Create a function that prints the multiplication table for a given number.
func multiplicationTable(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", number, i, number*i)
	}
}

// question 5 - Write a function that counts the number of digits in an integer.
func countDigit(number int) {

	if number == 0 {
		fmt.Println("The number of is 1")
	}

	if number < 0 {
		number = -number
	}

	count := 0
	for number > 0 {
		number = number / 10
		count++
	}

	fmt.Printf("The number of digits is %d", count)

}

func main() {
	fmt.Println("Hello From Loops")
	// sumOfN(10)
	// factorial(10)
	// fib(10)
	// multiplicationTable(9)
	countDigit(123456789)
}
