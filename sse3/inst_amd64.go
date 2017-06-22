package sse3

// go:noescape

// ADDSUBPDm128byte Packed Double-FP Add/Subtract
func ADDSUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// ADDSUBPDm128float64 Packed Double-FP Add/Subtract
func ADDSUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// ADDSUBPSm128byte Packed Single-FP Add/Subtract
func ADDSUBPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// ADDSUBPSm128float32 Packed Single-FP Add/Subtract
func ADDSUBPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// HADDPDm128byte Packed Double-FP Horizontal Add
func HADDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// HADDPDm128float64 Packed Double-FP Horizontal Add
func HADDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// HADDPSm128byte Packed Single-FP Horizontal Add
func HADDPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// HADDPSm128float32 Packed Single-FP Horizontal Add
func HADDPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// HSUBPDm128byte Packed Double-FP Horizontal Subtract
func HSUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// HSUBPDm128float64 Packed Double-FP Horizontal Subtract
func HSUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// HSUBPSm128byte Packed Single-FP Horizontal Subtract
func HSUBPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// HSUBPSm128float32 Packed Single-FP Horizontal Subtract
func HSUBPSm128float32(X1 []float32, X2 []float32)
