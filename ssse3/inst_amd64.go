package ssse3



// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSBm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSBm128uint8(X1 []uint8,  X2 []uint8)


// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSDm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSDm128uint32(X1 []uint32,  X2 []uint32)


// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSWm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Packed Absolute Value Unsigned Integers
func PABSWm128uint16(X1 []uint16,  X2 []uint16)


// go:noescape

// Packed Horizontal Add
func PHADDDm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Add and Saturate
func PHADDSWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Add
func PHADDWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Subtract
func PHSUBDm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Subtract and Saturate
func PHSUBSWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Subtract
func PHSUBWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Multiply and Add Packed Signed and Unsigned Bytes
func PMADDUBSWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Multiply High with Round and Scale
func PMULHRSWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Shuffle Bytes
func PSHUFBm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed SIGN
func PSIGNBm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed SIGN
func PSIGNDm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed SIGN
func PSIGNWm128byte(X1 []byte,  X2 []byte)

