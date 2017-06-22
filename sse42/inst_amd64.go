package sse42

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTQm128int64(X1 []int64, X2 []int64)
