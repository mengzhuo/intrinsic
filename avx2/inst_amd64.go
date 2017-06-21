package avx2



// go:noescape

// Load with Broadcast Floating-Point Data
func VBROADCASTSSbyte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Absolute Value
func VPABSBm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Packed Absolute Value
func VPABSDm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Packed Absolute Value
func VPABSWm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Load with Broadcast Integer Data from General Purpose Register
func VPBROADCASTBm8byte(X1 []byte,  X2 []byte)


// go:noescape

// Load with Broadcast Integer Data from General Purpose Register
func VPBROADCASTDm32byte(X1 []byte,  X2 []byte)


// go:noescape

// Load with Broadcast Integer Data from General Purpose Register
func VPBROADCASTQm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Load with Broadcast Integer Data from General Purpose Register
func VPBROADCASTWm16byte(X1 []byte,  X2 []byte)


// go:noescape

// Variable Bit Shift Left Logical
func VPSLLVDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Variable Bit Shift Left Logical
func VPSLLVQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Variable Bit Shift Right Arithmetic
func VPSRAVDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Variable Bit Shift Right Logical
func VPSRLVDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Variable Bit Shift Right Logical
func VPSRLVQm128byte(X1 []byte,  X2 []byte,  X3 []byte)

