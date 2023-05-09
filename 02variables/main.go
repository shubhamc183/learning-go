package main

import "fmt"

/**
CANNOT defined variables using Walrus operator outside function
numberOfDormantUser := 1000
*/

// Capital first letter represents the public keyword in go and is accesible anywhere in the file
const MauInMillions int = 180

func main() {

	var username string = "jessePinkman"
	fmt.Printf("var is of type: %T\n", username)

	var isUserLoggedIn bool = true
	fmt.Printf("var is of type: %T\n", isUserLoggedIn)

	var number uint8 = 44
	fmt.Printf("var is of type: %T\n", number)

	var value float32 = 231.4523
	fmt.Printf("var is of type: %T\n", value)

	// default value
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("var is of type: %T\n", anotherVar)

	var someString string
	fmt.Println(someString == "")
	fmt.Printf("var is of type: %T\n", someString)

	// implicit type
	var myData = "breaking bad"
	fmt.Println(myData)
	fmt.Printf("var is of type: %T\n", myData)
	myData = "better call saul"
	// myData = 120 can't change the data type of implict type variables

	// walrus operator / no var style
	numberOfUser := 100000000
	fmt.Println(numberOfUser)
	fmt.Printf("var is of type: %T\n", myData)

	fmt.Println(MauInMillions)
	fmt.Printf("var is of type: %T\n", MauInMillions)
}
