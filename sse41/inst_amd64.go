package sse41

// go:noescape

// Pack with Unsigned Saturation
func PACKUSDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Qword Data for Equal
func PCMPEQQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Horizontal Word Minimum
func PHMINPOSUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func PMAXSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func PMAXSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Maximum of Packed Signed Integers
func PMAXSDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func PMAXSDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Maximum of Packed Unsigned Integers
func PMAXUDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Unsigned Integers
func PMAXUDm128uint32(X1 []uint32, X2 []uint32)

// go:noescape

// Maximum of Packed Unsigned Integers
func PMAXUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Unsigned Integers
func PMAXUWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// Minimum of Packed Signed Integers
func PMINSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Signed Integers
func PMINSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Minimum of Packed Signed Integers
func PMINSDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Signed Integers
func PMINSDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Minimum of Packed Unsigned Integers
func PMINUDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Unsigned Integers
func PMINUDm128uint32(X1 []uint32, X2 []uint32)

// go:noescape

// Minimum of Packed Unsigned Integers
func PMINUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Unsigned Integers
func PMINUWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXBDm32byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXBQm16byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXBWm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXDQm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXWDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Sign Extend
func PMOVSXWQm32byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXBDm32byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXBQm16byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXBWm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXDQm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXWDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Packed Move with Zero Extend
func PMOVZXWQm32byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Doubleword Integers
func PMULDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Doubleword Integers
func PMULDQm128int64(X1 []int64, X2 []int64)

// go:noescape

// Multiply Packed Integers and Store Low Result
func PMULLDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Integers and Store Low Result
func PMULLDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PTEST- Logical Compare
func PTESTm128byte(X1 []byte, X2 []byte)
