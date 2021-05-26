package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("What program to run?: ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	switch x {
	case 1:
		mainSequential()
	case 2:
		main2()
	default:
		fmt.Println("No such case")
	}

}
