package ssse3

// go:noescape

// Packed Absolute Value Integers
func PABSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Absolute Value Integers
func PABSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Packed Absolute Value Integers
func PABSDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Absolute Value Integers
func PABSDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Packed Absolute Value Integers
func PABSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Absolute Value Integers
func PABSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Packed Horizontal Add
func PHADDDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Add and Saturate
func PHADDSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Add
func PHADDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Subtract
func PHSUBDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Subtract and Saturate
func PHSUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Subtract
func PHSUBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply and Add Packed Signed and Unsigned Bytes
func PMADDUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Multiply High with Round and Scale
func PMULHRSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Shuffle Bytes
func PSHUFBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed SIGN
func PSIGNBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed SIGN
func PSIGNDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed SIGN
func PSIGNWm128byte(X1 []byte, X2 []byte)
