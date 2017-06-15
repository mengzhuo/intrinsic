package sse3

import "testing"

func TestHADDPSm128float32(t *testing.T) {
	a := []float32{1, 2, 3, 4}
	b := []float32{4, 3, 2, 1}
	HADDPSm128float32(a, b)
	t.Log(a, b)
}
