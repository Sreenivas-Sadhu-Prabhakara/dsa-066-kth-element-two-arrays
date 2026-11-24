package main

import (
	"fmt"
	"math"
)

// KthElement returns the k-th smallest (1-based) element of the merged
// order of two sorted arrays, without merging. ok is false when k is out
// of range. Same partition search as the two-array median, aimed at k.
func KthElement(a, b []int, k int) (int, bool) {
	if len(a) > len(b) {
		a, b = b, a
	}
	n, m := len(a), len(b)
	if k < 1 || k > n+m {
		return 0, false
	}
	lo, hi := max(0, k-m), min(k, n)
	for lo <= hi {
		i := (lo + hi) / 2
		j := k - i
		aLeft, aRight := math.MinInt, math.MaxInt
		if i > 0 {
			aLeft = a[i-1]
		}
		if i < n {
			aRight = a[i]
		}
		bLeft, bRight := math.MinInt, math.MaxInt
		if j > 0 {
			bLeft = b[j-1]
		}
		if j < m {
			bRight = b[j]
		}
		switch {
		case aLeft <= bRight && bLeft <= aRight:
			return max(aLeft, bLeft), true
		case aLeft > bRight:
			hi = i - 1
		default:
			lo = i + 1
		}
	}
	return 0, false // unreachable for valid sorted inputs
}

func main() {
	v, _ := KthElement([]int{2, 3, 6, 7, 9}, []int{1, 4, 8, 10}, 5)
	fmt.Println(v)
}
