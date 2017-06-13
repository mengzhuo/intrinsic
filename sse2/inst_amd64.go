package sse2



// Add Packed Double-Precision Floating-Point Values
// go:noescape
func ADDPDm128byte(X1 []byte,  X2 []byte)

// Add Packed Double-Precision Floating-Point Values
// go:noescape
func ADDPDm128float32(X1 []float32,  X2 []float32)


// Add Scalar Double-Precision Floating-Point Values
// go:noescape
func ADDSDm64byte(X1 []byte,  X2 []byte)

// Add Scalar Double-Precision Floating-Point Values
// go:noescape
func ADDSDm64float32(X1 []float32,  X2 []float32)


// Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values
// go:noescape
func ANDNPDm128byte(X1 []byte,  X2 []byte)

// Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values
// go:noescape
func ANDNPDm128float32(X1 []float32,  X2 []float32)


// Bitwise Logical AND of Packed Double-Precision Floating-Point Values
// go:noescape
func ANDPDm128byte(X1 []byte,  X2 []byte)

// Bitwise Logical AND of Packed Double-Precision Floating-Point Values
// go:noescape
func ANDPDm128float32(X1 []float32,  X2 []float32)


// Compare Scalar Ordered Double-Precision Floating- Point Values and Set EFLAGS
// go:noescape
func COMISDm64byte(X1 []byte,  X2 []byte)


// Divide Packed Double-Precision Floating-Point Values
// go:noescape
func DIVPDm128byte(X1 []byte,  X2 []byte)

// Divide Packed Double-Precision Floating-Point Values
// go:noescape
func DIVPDm128float32(X1 []float32,  X2 []float32)


// Divide Scalar Double-Precision Floating-Point Values
// go:noescape
func DIVSDm64byte(X1 []byte,  X2 []byte)

// Divide Scalar Double-Precision Floating-Point Values
// go:noescape
func DIVSDm64float32(X1 []float32,  X2 []float32)


// Store Selected Bytes of Double Quadword
// go:noescape
func MASKMOVDQUbyte(X1 []byte,  X2 []byte)


// Return Maximum Packed Double-Precision Floating- Point Values
// go:noescape
func MAXPDm128byte(X1 []byte,  X2 []byte)


// Return Maximum Scalar Double-Precision Floating-Point Value
// go:noescape
func MAXSDm64byte(X1 []byte,  X2 []byte)

// Return Maximum Scalar Double-Precision Floating-Point Value
// go:noescape
func MAXSDm64float32(X1 []float32,  X2 []float32)


// Return Minimum Packed Double-Precision Floating-Point Values
// go:noescape
func MINPDm128byte(X1 []byte,  X2 []byte)

// Return Minimum Packed Double-Precision Floating-Point Values
// go:noescape
func MINPDm128float32(X1 []float32,  X2 []float32)


// Return Minimum Scalar Double-Precision Floating-Point Value
// go:noescape
func MINSDm64byte(X1 []byte,  X2 []byte)

// Return Minimum Scalar Double-Precision Floating-Point Value
// go:noescape
func MINSDm64float32(X1 []float32,  X2 []float32)


// Multiply Packed Double-Precision Floating-Point Values
// go:noescape
func MULPDm128byte(X1 []byte,  X2 []byte)

// Multiply Packed Double-Precision Floating-Point Values
// go:noescape
func MULPDm128float32(X1 []float32,  X2 []float32)


// Multiply Scalar Double-Precision Floating-Point Values
// go:noescape
func MULSDm64byte(X1 []byte,  X2 []byte)

// Multiply Scalar Double-Precision Floating-Point Values
// go:noescape
func MULSDm64float32(X1 []float32,  X2 []float32)


// Bitwise Logical OR of Double-Precision Floating-Point Values
// go:noescape
func ORPDm128byte(X1 []byte,  X2 []byte)

// Bitwise Logical OR of Double-Precision Floating-Point Values
// go:noescape
func ORPDm128float32(X1 []float32,  X2 []float32)


// Pack with Signed Saturation
// go:noescape
func PACKSSDWm128byte(X1 []byte,  X2 []byte)


// Pack with Signed Saturation
// go:noescape
func PACKSSWBm128byte(X1 []byte,  X2 []byte)


// Pack with Unsigned Saturation
// go:noescape
func PACKUSWBm128byte(X1 []byte,  X2 []byte)


// Add Packed Integers
// go:noescape
func PADDBm128byte(X1 []byte,  X2 []byte)

// Add Packed Integers
// go:noescape
func PADDBm128int8(X1 []int8,  X2 []int8)


// Add Packed Integers
// go:noescape
func PADDDm128byte(X1 []byte,  X2 []byte)

// Add Packed Integers
// go:noescape
func PADDDm128int32(X1 []int32,  X2 []int32)


// Add Packed Quadword Integers
// go:noescape
func PADDQm128byte(X1 []byte,  X2 []byte)

// Add Packed Quadword Integers
// go:noescape
func PADDQm128int64(X1 []int64,  X2 []int64)


// Add Packed Signed Integers with Signed Saturation
// go:noescape
func PADDSBm128byte(X1 []byte,  X2 []byte)

// Add Packed Signed Integers with Signed Saturation
// go:noescape
func PADDSBm128int8(X1 []int8,  X2 []int8)


// Add Packed Signed Integers with Signed Saturation
// go:noescape
func PADDSWm128byte(X1 []byte,  X2 []byte)

// Add Packed Signed Integers with Signed Saturation
// go:noescape
func PADDSWm128int16(X1 []int16,  X2 []int16)


// Add Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PADDUSBm128byte(X1 []byte,  X2 []byte)

// Add Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PADDUSBm128uint8(X1 []uint8,  X2 []uint8)


// Add Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PADDUSWm128byte(X1 []byte,  X2 []byte)

// Add Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PADDUSWm128uint16(X1 []uint16,  X2 []uint16)


// Add Packed Integers
// go:noescape
func PADDWm128byte(X1 []byte,  X2 []byte)

// Add Packed Integers
// go:noescape
func PADDWm128int16(X1 []int16,  X2 []int16)


// Logical AND
// go:noescape
func PANDm128byte(X1 []byte,  X2 []byte)


// Logical AND NOT
// go:noescape
func PANDNm128byte(X1 []byte,  X2 []byte)


// Average Packed Integers
// go:noescape
func PAVGBm128byte(X1 []byte,  X2 []byte)

// Average Packed Integers
// go:noescape
func PAVGBm128int8(X1 []int8,  X2 []int8)


// Average Packed Integers
// go:noescape
func PAVGWm128byte(X1 []byte,  X2 []byte)

// Average Packed Integers
// go:noescape
func PAVGWm128int16(X1 []int16,  X2 []int16)


// Compare Packed Data for Equal
// go:noescape
func PCMPEQBm128byte(X1 []byte,  X2 []byte)


// Compare Packed Data for Equal
// go:noescape
func PCMPEQDm128byte(X1 []byte,  X2 []byte)


// Compare Packed Data for Equal
// go:noescape
func PCMPEQWm128byte(X1 []byte,  X2 []byte)


// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTBm128byte(X1 []byte,  X2 []byte)

// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTBm128int8(X1 []int8,  X2 []int8)


// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTDm128byte(X1 []byte,  X2 []byte)

// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTDm128int32(X1 []int32,  X2 []int32)


// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTWm128byte(X1 []byte,  X2 []byte)

// Compare Packed Signed Integers for Greater Than
// go:noescape
func PCMPGTWm128int16(X1 []int16,  X2 []int16)


// Multiply and Add Packed Integers
// go:noescape
func PMADDWDm128byte(X1 []byte,  X2 []byte)

// Multiply and Add Packed Integers
// go:noescape
func PMADDWDm128int32(X1 []int32,  X2 []int32)


// Maximum of Packed Signed Word Integers
// go:noescape
func PMAXSWm128byte(X1 []byte,  X2 []byte)

// Maximum of Packed Signed Word Integers
// go:noescape
func PMAXSWm128int16(X1 []int16,  X2 []int16)


// Maximum of Packed Unsigned Byte Integers
// go:noescape
func PMAXUBm128byte(X1 []byte,  X2 []byte)

// Maximum of Packed Unsigned Byte Integers
// go:noescape
func PMAXUBm128uint8(X1 []uint8,  X2 []uint8)


// Minimum of Packed Signed Word Integers
// go:noescape
func PMINSWm128byte(X1 []byte,  X2 []byte)

// Minimum of Packed Signed Word Integers
// go:noescape
func PMINSWm128int16(X1 []int16,  X2 []int16)


// Minimum of Packed Unsigned Byte Integers
// go:noescape
func PMINUBm128byte(X1 []byte,  X2 []byte)

// Minimum of Packed Unsigned Byte Integers
// go:noescape
func PMINUBm128uint8(X1 []uint8,  X2 []uint8)


// Multiply Packed Unsigned Integers and Store High Result
// go:noescape
func PMULHUWm128byte(X1 []byte,  X2 []byte)

// Multiply Packed Unsigned Integers and Store High Result
// go:noescape
func PMULHUWm128uint16(X1 []uint16,  X2 []uint16)


// Multiply Packed Signed Integers and Store High Result
// go:noescape
func PMULHWm128byte(X1 []byte,  X2 []byte)

// Multiply Packed Signed Integers and Store High Result
// go:noescape
func PMULHWm128int16(X1 []int16,  X2 []int16)


// Multiply Packed Signed Integers and Store Low Result
// go:noescape
func PMULLWm128byte(X1 []byte,  X2 []byte)

// Multiply Packed Signed Integers and Store Low Result
// go:noescape
func PMULLWm128int16(X1 []int16,  X2 []int16)


// Multiply Packed Unsigned Doubleword Integers
// go:noescape
func PMULUDQm128byte(X1 []byte,  X2 []byte)

// Multiply Packed Unsigned Doubleword Integers
// go:noescape
func PMULUDQm128int64(X1 []int64,  X2 []int64)


// Bitwise Logical OR
// go:noescape
func PORm128byte(X1 []byte,  X2 []byte)


// Compute Sum of Absolute Differences
// go:noescape
func PSADBWm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Left Logical
// go:noescape
func PSLLDm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Left Logical
// go:noescape
func PSLLQm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Left Logical
// go:noescape
func PSLLWm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Right Arithmetic
// go:noescape
func PSRADm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Right Arithmetic
// go:noescape
func PSRAWm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Right Logical
// go:noescape
func PSRLDm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Right Logical
// go:noescape
func PSRLQm128byte(X1 []byte,  X2 []byte)


// Shift Packed Data Right Logical
// go:noescape
func PSRLWm128byte(X1 []byte,  X2 []byte)


// Subtract Packed Integers
// go:noescape
func PSUBBm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Integers
// go:noescape
func PSUBBm128int8(X1 []int8,  X2 []int8)


// Subtract Packed Integers
// go:noescape
func PSUBDm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Integers
// go:noescape
func PSUBDm128int32(X1 []int32,  X2 []int32)


// Subtract Packed Quadword Integers
// go:noescape
func PSUBQm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Quadword Integers
// go:noescape
func PSUBQm128int64(X1 []int64,  X2 []int64)


// Subtract Packed Signed Integers with Signed Saturation
// go:noescape
func PSUBSBm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Signed Integers with Signed Saturation
// go:noescape
func PSUBSBm128int8(X1 []int8,  X2 []int8)


// Subtract Packed Signed Integers with Signed Saturation
// go:noescape
func PSUBSWm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Signed Integers with Signed Saturation
// go:noescape
func PSUBSWm128int16(X1 []int16,  X2 []int16)


// Subtract Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PSUBUSBm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PSUBUSBm128uint8(X1 []uint8,  X2 []uint8)


// Subtract Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PSUBUSWm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Unsigned Integers with Unsigned Saturation
// go:noescape
func PSUBUSWm128uint16(X1 []uint16,  X2 []uint16)


// Subtract Packed Integers
// go:noescape
func PSUBWm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Integers
// go:noescape
func PSUBWm128int16(X1 []int16,  X2 []int16)


// Unpack High Data
// go:noescape
func PUNPCKHBWm128byte(X1 []byte,  X2 []byte)


// Unpack High Data
// go:noescape
func PUNPCKHDQm128byte(X1 []byte,  X2 []byte)


// Unpack High Data
// go:noescape
func PUNPCKHQDQm128byte(X1 []byte,  X2 []byte)


// Unpack High Data
// go:noescape
func PUNPCKHWDm128byte(X1 []byte,  X2 []byte)


// Unpack Low Data
// go:noescape
func PUNPCKLBWm128byte(X1 []byte,  X2 []byte)


// Unpack Low Data
// go:noescape
func PUNPCKLDQm128byte(X1 []byte,  X2 []byte)


// Unpack Low Data
// go:noescape
func PUNPCKLQDQm128byte(X1 []byte,  X2 []byte)


// Unpack Low Data
// go:noescape
func PUNPCKLWDm128byte(X1 []byte,  X2 []byte)


// Logical Exclusive OR
// go:noescape
func PXORm128byte(X1 []byte,  X2 []byte)


// Compute Square Roots of Packed Double-Precision Floating-Point Values
// go:noescape
func SQRTPDm128byte(X1 []byte,  X2 []byte)

// Compute Square Roots of Packed Double-Precision Floating-Point Values
// go:noescape
func SQRTPDm128float32(X1 []float32,  X2 []float32)


// Compute Square Root of Scalar Double-Precision Floating-Point Value
// go:noescape
func SQRTSDm64byte(X1 []byte,  X2 []byte)

// Compute Square Root of Scalar Double-Precision Floating-Point Value
// go:noescape
func SQRTSDm64float32(X1 []float32,  X2 []float32)


// Subtract Packed Double-Precision Floating-Point Values
// go:noescape
func SUBPDm128byte(X1 []byte,  X2 []byte)

// Subtract Packed Double-Precision Floating-Point Values
// go:noescape
func SUBPDm128float32(X1 []float32,  X2 []float32)


// Subtract Scalar Double-Precision Floating-Point Values
// go:noescape
func SUBSDm64byte(X1 []byte,  X2 []byte)

// Subtract Scalar Double-Precision Floating-Point Values
// go:noescape
func SUBSDm64float32(X1 []float32,  X2 []float32)


// Unordered Compare Scalar Double-Precision Floating- Point Values and Set EFLAGS
// go:noescape
func UCOMISDm64byte(X1 []byte,  X2 []byte)


// Unpack and Interleave High Packed Double- Precision Floating-Point Values
// go:noescape
func UNPCKHPDm128byte(X1 []byte,  X2 []byte)


// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
// go:noescape
func UNPCKLPDm128byte(X1 []byte,  X2 []byte)

// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
// go:noescape
func UNPCKLPDm128float32(X1 []float32,  X2 []float32)


// Bitwise Logical XOR for Double-Precision Floating-Point Values
// go:noescape
func XORPDm128byte(X1 []byte,  X2 []byte)

// Bitwise Logical XOR for Double-Precision Floating-Point Values
// go:noescape
func XORPDm128float32(X1 []float32,  X2 []float32)

