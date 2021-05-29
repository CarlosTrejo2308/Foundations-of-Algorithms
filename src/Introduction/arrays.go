package main

func AddArray(n int, S []int) (result int) {
	result = 0

	for i := 0; i < n; i++ {
		result += S[i]
	}

	return
}

func ExchangeSort(n int, S []int) {

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if S[j] < S[i] {
				S[i], S[j] = S[j], S[i]
			}
		}
	}

}

func MatrixMultiplication(n int, A, B, C [][]int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j]
			}
		}
	}
}
