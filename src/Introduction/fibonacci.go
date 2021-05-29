package main

import (
	"fmt"
	"strconv"
)

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	} else {
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}
}

func FibonacciIterative(n int) int {
	f := make([]int, n+1)

	f[0] = 0

	if n > 0 {
		f[1] = 1
		for i := 2; i <= n; i++ {
			f[i] = f[i-1] + f[i-2]
		}
	}

	return f[n]

}

func mainFibo(recursive bool) {

	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	n, _ := strconv.Atoi(userInput)

	var result int
	if recursive {
		result = FibonacciRecursive(n)
	} else {
		result = FibonacciIterative(n)
	}

	fmt.Printf("Rec(%d): %d\n", n, result)

}
