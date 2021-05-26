package main

import (
	"bufio"
	"os"
	"strings"
)

// Get input form the user
func GetUserData() string {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text

}
