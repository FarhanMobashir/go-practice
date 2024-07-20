package main

import (
	"fmt"
	"sort"
	"strings"
)

// question 1 - Write a function that counts the frequency of each word in a given text.
func countFrequency(text string) {
	text = strings.ToLower(text)

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// question 2 - Create a simple phone book using a map, allowing adding, deleting, and looking up entries.
type PhoneBook map[string]string

func initPhoneBook() {
	phoneBook := make(PhoneBook)

	for {
		var choice int
		fmt.Println("\nPhone book menu:")
		fmt.Println("1. Add Entry")
		fmt.Println("2. Delete Entry")
		fmt.Println("3. Look up Entry")
		fmt.Println("4. Exit")
		fmt.Println("\nEnter your choice below")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEntry(phoneBook)
		case 2:
			deleteEntry(phoneBook)
		case 3:
			lookUpEntry(phoneBook)
		case 4:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}

func addEntry(pb PhoneBook) {
	var name, number string
	fmt.Println("Enter name:")
	fmt.Scanln(&name)
	fmt.Println("Enter number:")
	fmt.Scanln(&number)
	pb[name] = number
	fmt.Println("Entry added")
}

func deleteEntry(pb PhoneBook) {
	var name string
	fmt.Print("Enter name to delete: ")
	fmt.Scanln(&name)
	if _, exists := pb[name]; exists {
		delete(pb, name)
		fmt.Println("Entry deleted.")
	} else {
		fmt.Println("Entry not found.")
	}
}

func lookUpEntry(pb PhoneBook) {
	var name string
	fmt.Print("Enter name to look up: ")
	fmt.Scanln(&name)
	if number, exists := pb[name]; exists {
		fmt.Printf("Name: %s, Phone Number: %s\n", name, number)
	} else {
		fmt.Println("Entry not found.")
	}
}

// question 3 - Implement a map with struct values and write functions to manipulate it. [Library Management]
type Book struct {
	Title           string
	Author          string
	ISBN            string
	AvailableCopies int
}

type Library map[string]Book

func initLibrary() {
	library := make(Library)

	for {
		var choice int
		fmt.Println("\nLibrary Management System Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Get Book Details")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook(library)
		case 2:
			removeBook(library)
		case 3:
			getBookDetails(library)
		case 4:
			borrowBook(library)
		case 5:
			returnBook(library)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addBook(lib Library) {
	var title, author, isbn string
	var copies int

	fmt.Println("Enter title: ")
	fmt.Scanln(&title)
	fmt.Println("Enter author: ")
	fmt.Scanln(&author)
	fmt.Println("Enter ISBN: ")
	fmt.Scanln(&isbn)
	fmt.Println("Enter number of available copies: ")
	fmt.Scanln(&copies)
	lib[isbn] = Book{Title: title, Author: author, ISBN: isbn, AvailableCopies: copies}
	fmt.Println("Book added.")
}

func removeBook(lib Library) {
	var isbn string
	fmt.Print("Enter ISBN to remove: ")
	fmt.Scanln(&isbn)
	if _, exists := lib[isbn]; exists {
		delete(lib, isbn)
		fmt.Println("Book removed.")
	} else {
		fmt.Println("Book not found.")
	}
}

func getBookDetails(lib Library) {
	var isbn string
	fmt.Print("Enter ISBN to get details: ")
	fmt.Scanln(&isbn)
	if book, exists := lib[isbn]; exists {
		fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available Copies: %d\n", book.Title, book.Author, book.ISBN, book.AvailableCopies)
	} else {
		fmt.Println("Book not found.")
	}
}

func borrowBook(lib Library) {
	var isbn string
	fmt.Print("Enter ISBN to borrow: ")
	fmt.Scanln(&isbn)
	if book, exists := lib[isbn]; exists {
		if book.AvailableCopies > 0 {
			book.AvailableCopies--
			lib[isbn] = book
			fmt.Println("Book borrowed.")
		} else {
			fmt.Println("No available copies.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func returnBook(lib Library) {
	var isbn string
	fmt.Print("Enter ISBN to return: ")
	fmt.Scanln(&isbn)
	if book, exists := lib[isbn]; exists {
		book.AvailableCopies++
		lib[isbn] = book
		fmt.Println("Book returned.")
	} else {
		fmt.Println("Book not found.")
	}
}

// question 4 - Write a function to invert the keys and values of a map.
func invertMap(m map[string]string) {
	fmt.Println("Before :", m)
	for key, value := range m {
		tempKey := key
		tempValue := value
		delete(m, tempKey)
		m[tempValue] = tempKey
	}

	fmt.Println("After", m)
}

// question 5 - Implement a function to sort a map by its keys.
func sortMapByKeys(m map[string]int) {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, m[key])
	}
}

func main() {
	// text := "Hello world! Hello everyone. This is a test. Hello, Hello, world."
	// countFrequency(text)
	// initPhoneBook()
	// initLibrary()

	invertMap(map[string]string{
		"Alice":   "123-456-7890",
		"Bob":     "987-654-3210",
		"Charlie": "555-555-5555",
	})

	sortMapByKeys(map[string]int{"orange": 10, "kivi": 4, "apple": 5, "banana": 3, "cherry": 2})
}
