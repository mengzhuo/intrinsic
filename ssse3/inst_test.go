package ssse3

import (
	"bytes"
	"testing"
)

func TestPSHUFBm128byte(t *testing.T) {
	a := bytes.Repeat([]byte{2}, 16)
	b := bytes.Repeat([]byte{1}, 16)
	PSHUFBm128byte(a, b)
	t.Log(a, b)
}
