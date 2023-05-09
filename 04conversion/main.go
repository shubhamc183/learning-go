package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our store")
	fmt.Println("Please rate our pizza")
	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to the input rating: ", numRating+1)
	}
}
