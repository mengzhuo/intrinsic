package sse3

// Packed Double-FP Add/Subtract
// go:noescape
func ADDSUBPDm128byte(X1 []byte, X2 []byte)

// Packed Single-FP Add/Subtract
// go:noescape
func ADDSUBPSm128byte(X1 []byte, X2 []byte)

// Packed Double-FP Horizontal Add
// go:noescape
func HADDPDm128byte(X1 []byte, X2 []byte)

// Packed Single-FP Horizontal Add
// go:noescape
func HADDPSm128byte(X1 []byte, X2 []byte)

// Packed Double-FP Horizontal Subtract
// go:noescape
func HSUBPDm128byte(X1 []byte, X2 []byte)

// Packed Single-FP Horizontal Subtract
// go:noescape
func HSUBPSm128byte(X1 []byte, X2 []byte)
