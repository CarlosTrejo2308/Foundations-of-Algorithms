package main

import (
	"fmt"
	"math"
	"strconv"
)

func BinSearch(S []int, x int, location *int) int {

	comparisons := 0

	high := len(S)

	var low, mid int

	*location = -1

	for low <= high && *location == -1 {
		comparisons++
		mid = int(math.Floor(float64(low+high) / float64(2.0)))

		// Basic operation
		if x == S[mid] {
			*location = mid
		} else if x < S[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return comparisons
}

var array = []int{3, 4, 8, 9, 12, 24, 27, 70, 80, 81, 82, 83, 90, 100}

func mainBinSearch() {
	var location int

	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	// Search
	comp := BinSearch(array, x, &location)

	// Print the result
	if location < 0 {
		fmt.Printf("The number %d was not in the list", x)
	} else {
		fmt.Printf("The number %d is in the index %d", x, location)
	}
	fmt.Printf(" with %d comparisions\n", comp)

}

// Searchs in the whole array to find an element
// n, lenght of the array, input size
// S, array to search
// x, number to find
// location, index to save the index

// W(n) = n

/* Average
k: array slot

if x in the array
A(n) = (n + 1) / n


if x may not be in the array
A(n) = n * ( 1 - ( p / n ) ) + ( p / 2 )
*/
func SequentialSearch(n int, S []int, x int, location *int) int {
	comparisons := 0

	*location = 0

	// Keep looking until we find the elemnt or we run the whole array
	for *location < n && S[*location] != x { // Basic operation
		comparisons++
		*location++
	}

	// We didn't find it, thus we return a negative value
	if *location >= n {
		*location = -1
	}
	return comparisons
}

func mainSequential() {
	n := cap(array)
	var location int

	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	// Search
	comp := SequentialSearch(n, array, x, &location)

	// Print the result
	if location < 0 {
		fmt.Printf("The number %d was not in the list", x)
	} else {
		fmt.Printf("The number %d is in the index %d", x, location)
	}
	fmt.Printf(" with %d comparisions\n", comp)

}
