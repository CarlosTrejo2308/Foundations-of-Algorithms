package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Choose a program to run:\n\t1. Sequential Search\n\t2. Add array members\n\t3. Exchange sort")
	fmt.Print("> ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	switch x {
	case 1:
		mainSequential()
	case 2:
		temp := addArray( 5, []int{1, 2, 3, 20, -1} )
		fmt.Println(temp)
	default:
		fmt.Println("No such case")
	}

}
