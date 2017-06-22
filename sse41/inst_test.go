package sse41

import "testing"

func TestPMAXSBm128int8(t *testing.T) {
	a := []int8{1, 3, 4, 5, 6, 7, -1, -2, 1, 3, 4, 5, 6, 7, -1}
	b := []int8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	PMAXSBm128int8(a, b)
	t.Log(a, b)
}
