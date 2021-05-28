package main

func addArray(n int, S []int) (result int) {
	result = 0

	for i := 0; i < n; i++ {
		result += S[i]
	}

	return
}

func exchangeSort(n int, S []int) {

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if S[j] < S[i] {
				S[i], S[j] = S[j], S[i]
			}
		}
	}

}
