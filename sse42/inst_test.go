package sse42

import "testing"

func TestPCMPGTQm128byte(t *testing.T) {

	a := []byte{1, 3, 4, 5, 6, 7, 1, 2, 1, 3, 4, 5, 6, 7, 1}
	b := []byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	PCMPGTQm128byte(a, b)
	t.Log(a, b)
}
