package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango"}
	fmt.Println("fruitList is", fruitList)
	fmt.Println("fruitList's length is", len(fruitList))
	fmt.Printf("fruitList's type is %T\n", fruitList)

	fruitList = append(fruitList, "Banana")
	fruitList = append(fruitList, "Peach", "Watermelon")
	fmt.Println("fruitList is", fruitList)
	fmt.Println("fruitList's length is", len(fruitList))

	fruitList = append(fruitList[1:])
	fmt.Println("fruitList is", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruitList is", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 213123
	highScores[2] = 8412
	highScores[3] = 23

	fmt.Println(highScores)

	// this would give error as the slice is created of size 4
	// highScores[4] = 12332

	highScores = append(highScores, 1322113)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
