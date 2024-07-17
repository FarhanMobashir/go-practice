package main

import "fmt"

// question 1 - Write  program that prints "Hello,Golang!" to the console
func sayHello() string {
  return "Hello, Golang!"
}

// question 2 - Create a program that declares a variable and a constant , then print their values
func printVariableAndConstant() {
    var name string = "Mobashir"
    const age = 21
    fmt.Printf("My name is %s, and age is %d", name,age)
}

// question 3 - Write a function that takes two integer and return their sum, diff, product, and 
              // quotient
func basicArithmetic(a int, b int) {
  sum := a +b
  diff := a -b
  prod := a*b
  quotient := a % b

  fmt.Printf("sum -> %d, product -> %d, diff -> %d, quotient -> %d ", sum,prod,diff,quotient)
}

// question 4 - Implement a function to convert temperature between Celcius and Fahrenheit
func celciusToFahrenheit(temperature float32) {
  converted := float32((9.0/5.0 * temperature) + 32)
  
  fmt.Printf("The temperature in Fahrenheit %.2f", converted)
}

// question 5 - Write a program that calculates simple interest given principal, rate and time
func simpleInterest(p float64, r float64,t float64)  {
  si := (p * r * t)/100

  fmt.Printf("The Simple Interest is %f", si)
}

func main() {
  // fmt.Println(sayHello())
  // printVariableAndConstant()
  // basicArithmetic(10,5)
  // celciusToFahrenheit(-40)
  simpleInterest(1000, 5, 2)
} 
