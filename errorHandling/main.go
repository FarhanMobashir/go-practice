package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// question 1 - Demonstrate the use of panic and recover in a program.
func divide(a, b int) int {
	if b == 0 {
		panic("divison by zero")
	}
	return a / b
}

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result := divide(a, b)
	fmt.Printf("Result of %d / %d = %d\n", a, b, result)
}

// question 2 - Write a function that retries an operation a specified number of times if it fails.
func fetchDataWithRetry() ([]byte, error) {
	var data []byte
	var err error
	retries := 3

	for i := 0; i < retries; i++ {
		data, err = fetchData()

		if err == nil {
			return data, nil
		}
		fmt.Printf("Attempt %d failed: %v\n", i+1, err)
		time.Sleep(1 * time.Second)
	}
	return nil, fmt.Errorf("failed after %d attempts: %v", retries, err)
}

func fetchData() ([]byte, error) {
	if rand.Float32() < 0.7 { // Simulate 70% chance of failure
		return nil, errors.New("fetch failed")
	}
	return []byte("fetched data"), nil
}

func main() {
	fmt.Println("Start of main")
	// safeDivide(10, 2)
	// safeDivide(10, 0)

	data, err := fetchDataWithRetry()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}

	fmt.Println("End of main")
}
