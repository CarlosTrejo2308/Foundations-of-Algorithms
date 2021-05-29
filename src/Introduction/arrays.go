package main

// n: input size
// T(n) = n
func AddArray(n int, S []int) (result int) {
	result = 0

	for i := 0; i < n; i++ {
		result += S[i] // Basic operation
	}

	return
}

// Aka. Bubble sort
// n: Input size
// T(n) = (n-1) + (n-2) + ... + 1 = ( ( n - 1 ) * n ) / 2
func ExchangeSort(n int, S []int) {

	for i := 0; i < n; i++ {
		// 1: n-1 passes
		// 2: n-2 passes
		for j := i + 1; j < n; j++ {
			if S[j] < S[i] { // Basic operation
				S[i], S[j] = S[j], S[i]
			}
		}
	}

}

// n: Input size
// T(n) = n x n x n = n³
func MatrixMultiplication(n int, A, B, C [][]int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j] // Mult: Basic operation
			}
		}
	}
}

func Fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}
