package ssse3

// go:noescape

// PABSBm128byte Packed Absolute Value Integers
func PABSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PABSBm128int8 Packed Absolute Value Integers
func PABSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PABSDm128byte Packed Absolute Value Integers
func PABSDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PABSDm128int32 Packed Absolute Value Integers
func PABSDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PABSWm128byte Packed Absolute Value Integers
func PABSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PABSWm128int16 Packed Absolute Value Integers
func PABSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PHADDDm128byte Packed Horizontal Add
func PHADDDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PHADDSWm128byte Packed Horizontal Add and Saturate
func PHADDSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PHADDWm128byte Packed Horizontal Add
func PHADDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PHSUBDm128byte Packed Horizontal Subtract
func PHSUBDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PHSUBSWm128byte Packed Horizontal Subtract and Saturate
func PHSUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PHSUBWm128byte Packed Horizontal Subtract
func PHSUBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMADDUBSWm128byte Multiply and Add Packed Signed and Unsigned Bytes
func PMADDUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMULHRSWm128byte Packed Multiply High with Round and Scale
func PMULHRSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSHUFBm128byte Packed Shuffle Bytes
func PSHUFBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSIGNBm128byte Packed SIGN
func PSIGNBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSIGNDm128byte Packed SIGN
func PSIGNDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSIGNWm128byte Packed SIGN
func PSIGNWm128byte(X1 []byte, X2 []byte)
