package main

import "fmt"

func main() {
	// assigning the length of the array is the most important thing and also it's length would be 5
	// always irrespective of elements you fill inside
	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[0] = "Mango"
	fmt.Println("fruitList is", fruitList)
	fmt.Println("fruitList length is", len(fruitList))
	fmt.Printf("fruitList type is %T\n", fruitList)

	vegList := [5]string{"Carrot", "Tomato"}
	fmt.Println(vegList)
}
