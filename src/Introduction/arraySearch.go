package main

import (
	"fmt"
	"math"
	"strconv"
)

func BinSearch(S []int, x int, location *int) {
	high := len(S)

	var low, mid int

	*location = -1

	for low <= high && *location == -1 {
		mid = int(math.Floor(float64(low+high) / float64(2.0)))

		if x == S[mid] {
			*location = mid
		} else if x < S[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

func mainBinSearch() {
	// Default variables
	array := []int{3, 4, 8, 9, 12, 24, 27, 70, 80, 81, 82, 83, 90, 100}

	var location int

	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	// Search
	BinSearch(array, x, &location)

	// Print the result
	if location < 0 {
		fmt.Printf("The number %d was not in the list\n", x)
	} else {
		fmt.Printf("The number %d is in the index %d\n", x, location)
	}

}

// Searchs in the whole array to find an element
// n, lenght of the array
// S, array to search
// x, number to find
// location, index to save the index
func SequentialSearch(n int, S []int, x int, location *int) {
	*location = 0

	// Keep looking until we find the elemnt or we run the whole array
	for *location < n && S[*location] != x {
		*location++
	}

	// We didn't find it, thus we return a negative value
	if *location >= n {
		*location = -1
	}
}

func mainSequential() {
	// Default variables
	array := []int{3, 4, 8, 9, 12, 24, 27, 70}
	n := cap(array)
	var location int

	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	// Search
	SequentialSearch(n, array, x, &location)

	// Print the result
	if location < 0 {
		fmt.Printf("The number %d was not in the list\n", x)
	} else {
		fmt.Printf("The number %d is in the index %d\n", x, location)
	}

}
