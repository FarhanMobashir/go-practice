package main

import "fmt"

// question 1 - Write a function that takes an integer and returns "Even" if it's even and "Odd" if it's odd.
func evenOdd(a int) {
	if a%2 == 0 {
		fmt.Printf("Number %d is even\n", a)
	} else {
		fmt.Printf("Number %d is odd\n", a)
	}
}

// question 2 - Create a function that returns the maximum of three numbers.
func maxOfThree(a int, b int, c int) {

	if a > b && a > c {
		fmt.Printf("The maximum of three is %d\n", a)
	} else if b > a && b > c {
		fmt.Printf("The maximum of three is %d\n", b)
	} else {
		fmt.Printf("The maximum of three is %d\n", c)
	}

}

// question 3 - Write a function to check if a given year is a leap year.
func isLeapYear(year int) {
	if year%4 == 0 {
		fmt.Printf("The year %d is a Leap Year\n", year)
	} else {
		fmt.Printf("The year %d is not a Leap Year\n", year)
	}
}

// question 4 - Implement a program that takes a score and prints the corresponding grade (A, B, C, D, F).
func gradeCalculator(score int) {
	switch {
	case score >= 91 && score <= 100:
		fmt.Println("A")
	case score >= 81 && score <= 90:
		fmt.Println("B")
	case score >= 71 && score <= 80:
		fmt.Println("C")
	case score >= 61 && score <= 70:
		fmt.Println("D")
	case score >= 0 && score <= 60:
		fmt.Println("F")
	default:
		fmt.Println("Invalid score")
	}
}

// question 5 - Implement the FizzBuzz challenge
func fizzBuzz(limit int) {
	for i := 1; i <= limit; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// evenOdd(12)
	// evenOdd(17)
	// maxOfThree(3, 10, 5)
	// isLeapYear(2024)
	// gradeCalculator(55)
	fizzBuzz(100)
}
