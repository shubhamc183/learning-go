package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating")
	// read till
	input, _ := reader.ReadString('\n')

	// even \n is part of the input
	fmt.Println(input)
	fmt.Printf("The type of input is: %T\n", input)
}
