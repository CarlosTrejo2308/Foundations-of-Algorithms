package main

import (
	"fmt"
	"strconv"
)

/* // Get input form the user
func getUserData() string {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text

}
*/
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
		fmt.Printf("The number %d was not in the list", x)
	} else {
		fmt.Printf("The number %d is in the index %d", x, location)
	}

}
