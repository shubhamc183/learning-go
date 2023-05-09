package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	// this is how date formatting works in GoLang and we need to remember this in order to get the required format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1995, time.December, 12, 12, 20, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
