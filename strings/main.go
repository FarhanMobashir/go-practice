package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// question 1 - Write a function to check if a given string is a palindrome
func isPalindrome(str string) {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	if str == reversed {
		fmt.Printf("The string {%v} is a palindrome", str)
	} else {
		fmt.Printf("The string {%v} is not a palindrome", str)
	}
}

// question 2 - Create a function that counts the number of vowels in a given string.
func countVowels(word string) {
	vowelCount := 0

	for i := 0; i < len(word); i++ {
		letter := strings.ToLower(string(word[i]))
		if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" {
			vowelCount++
		}
	}
	fmt.Printf("The vowel count is %d", vowelCount)
}

// question 3 - Write a function to reverse a given string.
func reverseString(str string) {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	fmt.Printf("The reversed string is %v", reversed)
}

func sortString(s string) string {
	runes := []rune(strings.ToLower(s))

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

// question 4 - Implement a function to check if two string is anagrams.
func isAnagram(str1 string, str2 string) {
	if len(str1) != len(str2) {
		fmt.Println("Not anagrams")
		return
	}

	sorted1 := sortString(str1)
	sorted2 := sortString(str2)

	if sorted1 == sorted2 {
		fmt.Println("Is anagrams")
	} else {
		fmt.Println("Not anagrams")
	}

}

// question 5 - Write a function to compress a string using the counts of repeated characters.
func compressString(str string) {

	if len(str) == 0 {
		fmt.Println("The string is empty")
	}

	compressed := ""
	count := 1
	currentChar := rune(str[0])

	for i := 1; i < len(str); i++ {
		if rune(str[i]) == currentChar {
			count++
		} else {
			compressed += string(currentChar) + strconv.Itoa(count)
			currentChar = rune(str[i])
			count = 1
		}
	}

	compressed += string(currentChar) + strconv.Itoa(count)

	fmt.Printf("The compressed string is %v", compressed)
}

func main() {
	// fmt.Println("Hello from string")
	// isPalindrome("level")
	//countVowels("cauliflOwer")
	// reverseString("rihsabom")
	// isAnagram("Listen", "Silen")
	compressString("aaabcccccaaa")
}
