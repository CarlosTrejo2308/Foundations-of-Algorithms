package main

func addArray(n int, S []int) (result int) {
	result = 0

	for i := 0; i < n; i++ {
		result += S[i]
	}

	return
}
