package main

import (
	"fmt"
	"math"
)

// question 1 - Write a function that reverses an array.
func reverseArray(arr []int) {
	reverse := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}
	fmt.Println(reverse)
}

// question 2 - Create a function that returns the maximum value in an array.
func maxInArray(arr []int) {
	max := int(math.Inf(-1))
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	fmt.Printf("The max is %d", max)
}

// question 3 - Write a function that calculates the average of values in an array.
func averageOfArray(array []int) {

	if len(array) == 0 {
		fmt.Println("The array is empty, hence no average can be calculated")
	}

	sum := 0
	for _, value := range array {
		sum += value
	}

	average := float64(sum) / float64(len(array))
	fmt.Printf("The average of the array is %.1f", average)
}

// question 4 - Implement a function to merge two sorted arrays into a single sorted array.
func mergeSortedArray(array1 []int, array2 []int) {
	result := []int{}
	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(array1) && pointer2 < len(array2) {
		if array1[pointer1] < array2[pointer2] {
			result = append(result, array1[pointer1])
			pointer1++
		} else {
			result = append(result, array2[pointer2])
			pointer2++
		}
	}

	for pointer1 < len(array1) {
		result = append(result, array1[pointer1])
		pointer1++
	}
	for pointer2 < len(array2) {
		result = append(result, array2[pointer2])
		pointer2++
	}

	fmt.Println(result)

}

func isInArray(num int, arr []int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}

// question 5 - Write a function that removes duplicate values from an array.
func removeDuplicates(array []int) {
	removed := []int{}

	for i := 0; i < len(array); i++ {
		if !isInArray(array[i], removed) {
			removed = append(removed, array[i])
		}
	}

	fmt.Println(removed)
}

func main() {
	//reverseArray([]int{1, 2, 3, 4, 5})
	// maxInArray([]int{3, 5, 6, 10, -34})
	// averageOfArray([]int{1, 2, 3, 4, 5, 6})
	// mergeSortedArray([]int{1, 2, 3, 54, 67, 100}, []int{2, 3, 6, 8, 10})
	removeDuplicates([]int{1, 2, 2, 3, 3, 4, 5, 6, 7, 5, 7, 8, 9})
}
