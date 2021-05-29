package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Choose a program to run:\n\t1. Sequential Search\n\t2. Add array members\n\t3. Exchange sort")
	fmt.Println("\t4. Matrix multiplication\n\t5. Secuential Search\n\t6. Binary Search\n\t7. Fib Rec\n\t8. Fib Ite")

	fmt.Print("> ")
	userInput := GetUserData()
	x, _ := strconv.Atoi(userInput)

	switch x {
	case 1:
		mainSequential()
	case 2:
		temp := AddArray(5, []int{1, 2, 3, 20, -1})
		fmt.Println(temp)
	case 3:
		ejerciceArray := []int{4, 2, 6, 1, 0, 9}
		ExchangeSort(len(ejerciceArray), ejerciceArray)
		fmt.Println(ejerciceArray)
	case 4:
		ar1 := [][]int{{1, 3, 4}, {8, 3, 1}, {4, 1, 6}}
		ar2 := [][]int{{6, 4, 3}, {5, 3, 4}, {7, 8, 2}}
		n := 3
		ar3 := make([][]int, n)

		for i := range ar3 {
			ar3[i] = make([]int, n)
		}

		MatrixMultiplication(n, ar1, ar2, ar3)
		fmt.Println(ar3)

	case 5:
		mainSequential()
	case 6:
		mainBinSearch()
	case 7:
		mainFibo(true)
	case 8:
		mainFibo(false)
	default:
		fmt.Println("No such case")
	}

}
