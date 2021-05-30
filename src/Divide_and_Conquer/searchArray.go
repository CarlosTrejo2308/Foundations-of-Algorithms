package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var array = []int{3, 4, 8, 9, 12, 24, 27, 70, 80, 81, 82, 83, 90, 100}

// Get input form the user
// TODO: Move to independent module
func GetUserData() string {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text

}

func binSearchRecursive(n, low, high int, S []int) (location int) {
	var mid int

	if low > high {
		return -1
	} else {
		mid = int(math.Floor(float64(low+high) / 2.0))

		if S[mid] == n {
			return mid
		} else if n < S[mid] {
			return binSearchRecursive(n, low, mid-1, S)
		} else {
			return binSearchRecursive(n, mid+1, high, S)
		}
	}
}

// TODO: Move to independent module
func mainBinSearchRec() {
	// Get number to search
	fmt.Print("What number to find? : ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	// Search
	location := binSearchRecursive(x, 0, len(array), array)

	// Print the result
	if location < 0 {
		fmt.Printf("The number %d was not in the list\n", x)
	} else {
		fmt.Printf("The number %d is in the index %d\n", x, location)
	}

}
