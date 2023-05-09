package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers demo")

	// value is nil
	var ptr *int
	fmt.Println("Value of pointer is", ptr)

	count := 100
	// var countPtr *int = &count
	// even implicitly pointer can be defined
	countPtr := &count
	fmt.Println("Value of countPtr is", countPtr)
	fmt.Println("Value at countPtr is", *countPtr)

	// changing underlying data using pointers
	*countPtr = *countPtr + 10
	fmt.Println("Value of count is", count)
}
