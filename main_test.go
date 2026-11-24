package main

import "testing"

func TestKthElement(t *testing.T) {
	a := []int{2, 3, 6, 7, 9}
	b := []int{1, 4, 8, 10}
	// merged: 1 2 3 4 6 7 8 9 10
	wants := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	for k := 1; k <= 9; k++ {
		got, ok := KthElement(a, b, k)
		if !ok || got != wants[k-1] {
			t.Errorf("KthElement(k=%d) = (%d,%v), want %d", k, got, ok, wants[k-1])
		}
	}

	if _, ok := KthElement(a, b, 0); ok {
		t.Error("k=0 should be out of range")
	}
	if _, ok := KthElement(a, b, 10); ok {
		t.Error("k beyond total should be out of range")
	}
	if got, ok := KthElement([]int{}, []int{5}, 1); !ok || got != 5 {
		t.Errorf("single-element case: (%d,%v)", got, ok)
	}
}
