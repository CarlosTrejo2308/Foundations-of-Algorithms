package main

func merge(h, m int, U, V, S []int) {
	var i, j, k int

	for i < h && j < m {
		if U[i] < V[j] {
			S[k] = U[i]
			i++
		} else {
			S[k] = V[j]
			j++
		}

		k++
	}

	if i >= h {

	}
}
