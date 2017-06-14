package sse3

// go:noescape

// Packed Double-FP Add/Subtract
func ADDSUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Double-FP Add/Subtract
func ADDSUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Packed Single-FP Add/Subtract
func ADDSUBPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Single-FP Add/Subtract
func ADDSUBPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// Packed Double-FP Horizontal Add
func HADDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Double-FP Horizontal Add
func HADDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Packed Single-FP Horizontal Add
func HADDPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Single-FP Horizontal Add
func HADDPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// Packed Double-FP Horizontal Subtract
func HSUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Double-FP Horizontal Subtract
func HSUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Packed Single-FP Horizontal Subtract
func HSUBPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Single-FP Horizontal Subtract
func HSUBPSm128float32(X1 []float32, X2 []float32)
