package ssse3

import "testing"

func TestPABSB(t *testing.T) {
	a := []int8{2, 2, 0, 0, 3, -1, -2, 2, 2, 0, 0, 3, -1, -2, -9, -10}
	b := []int8{-1, -2, -3, -4, -5, -6, -7, -8, -1, -2, -3, -4, -5, -6, -7, -8}
	PABSBm128int8(a, b)
	t.Log(a, b)
}
