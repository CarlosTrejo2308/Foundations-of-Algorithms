package main

import "testing"

func TestAddArray(t *testing.T) {
	want := 15
	n := 5
	array1 := []int{1, 2, 3, 4, 5}

	got := AddArray(n, array1)

	if got != want {
		t.Errorf("addArray(n, {1, 2, 3, 4, 5}). Got %d Want %d", got, want)
	}
}
