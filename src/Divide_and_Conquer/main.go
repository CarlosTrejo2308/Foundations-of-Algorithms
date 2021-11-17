package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Choose a program to run:\n\t1. Binary Search (Recursive)")

	fmt.Print("> ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	switch x {
	case 1:
		mainBinSearchRec()
	default:
		fmt.Println("No such case")
	}

}
