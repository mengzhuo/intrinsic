package sse2

// go:noescape

// Add Packed Double-Precision Floating-Point Values
func ADDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Double-Precision Floating-Point Values
func ADDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Add Scalar Double-Precision Floating-Point Values
func ADDSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Add Scalar Double-Precision Floating-Point Values
func ADDSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values
func ANDNPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values
func ANDNPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Bitwise Logical AND of Packed Double-Precision Floating-Point Values
func ANDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Bitwise Logical AND of Packed Double-Precision Floating-Point Values
func ANDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Compare Scalar Ordered Double-Precision Floating- Point Values and Set EFLAGS
func COMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Divide Packed Double-Precision Floating-Point Values
func DIVPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Divide Packed Double-Precision Floating-Point Values
func DIVPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Divide Scalar Double-Precision Floating-Point Values
func DIVSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Divide Scalar Double-Precision Floating-Point Values
func DIVSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Store Selected Bytes of Double Quadword
func MASKMOVDQUbyte(X1 []byte, X2 []byte)

// go:noescape

// Return Maximum Packed Double-Precision Floating- Point Values
func MAXPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Return Maximum Scalar Double-Precision Floating-Point Value
func MAXSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Return Maximum Scalar Double-Precision Floating-Point Value
func MAXSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Return Minimum Packed Double-Precision Floating-Point Values
func MINPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Return Minimum Packed Double-Precision Floating-Point Values
func MINPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Return Minimum Scalar Double-Precision Floating-Point Value
func MINSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Return Minimum Scalar Double-Precision Floating-Point Value
func MINSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Multiply Packed Double-Precision Floating-Point Values
func MULPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Double-Precision Floating-Point Values
func MULPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Multiply Scalar Double-Precision Floating-Point Values
func MULSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Scalar Double-Precision Floating-Point Values
func MULSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Bitwise Logical OR of Double-Precision Floating-Point Values
func ORPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Bitwise Logical OR of Double-Precision Floating-Point Values
func ORPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Pack with Signed Saturation
func PACKSSDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Pack with Signed Saturation
func PACKSSWBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Pack with Unsigned Saturation
func PACKUSWBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Integers
func PADDBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Integers
func PADDBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Add Packed Integers
func PADDDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Integers
func PADDDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Add Packed Quadword Integers
func PADDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Quadword Integers
func PADDQm128int64(X1 []int64, X2 []int64)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func PADDSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func PADDSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func PADDSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func PADDSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// Add Packed Integers
func PADDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Add Packed Integers
func PADDWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Logical AND
func PANDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Logical AND NOT
func PANDNm128byte(X1 []byte, X2 []byte)

// go:noescape

// Average Packed Integers
func PAVGBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Average Packed Integers
func PAVGBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Average Packed Integers
func PAVGWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Average Packed Integers
func PAVGWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Compare Packed Data for Equal
func PCMPEQBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Data for Equal
func PCMPEQDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Data for Equal
func PCMPEQWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func PCMPGTWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Multiply and Add Packed Integers
func PMADDWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply and Add Packed Integers
func PMADDWDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Maximum of Packed Signed Word Integers
func PMAXSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Signed Word Integers
func PMAXSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Maximum of Packed Unsigned Byte Integers
func PMAXUBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Maximum of Packed Unsigned Byte Integers
func PMAXUBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// Minimum of Packed Signed Word Integers
func PMINSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Signed Word Integers
func PMINSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Minimum of Packed Unsigned Byte Integers
func PMINUBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Minimum of Packed Unsigned Byte Integers
func PMINUBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// Multiply Packed Signed Integers and Store High Result
func PMULHWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Signed Integers and Store High Result
func PMULHWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Multiply Packed Signed Integers and Store Low Result
func PMULLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Signed Integers and Store Low Result
func PMULLWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Multiply Packed Unsigned Doubleword Integers
func PMULUDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Multiply Packed Unsigned Doubleword Integers
func PMULUDQm128int64(X1 []int64, X2 []int64)

// go:noescape

// Bitwise Logical OR
func PORm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compute Sum of Absolute Differences
func PSADBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Left Logical
func PSLLDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Left Logical
func PSLLQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Left Logical
func PSLLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Right Arithmetic
func PSRADm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Right Arithmetic
func PSRAWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Right Logical
func PSRLDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Right Logical
func PSRLQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Shift Packed Data Right Logical
func PSRLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Integers
func PSUBBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Integers
func PSUBBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Subtract Packed Integers
func PSUBDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Integers
func PSUBDm128int32(X1 []int32, X2 []int32)

// go:noescape

// Subtract Packed Quadword Integers
func PSUBQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Quadword Integers
func PSUBQm128int64(X1 []int64, X2 []int64)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func PSUBSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func PSUBSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func PSUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func PSUBSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// Subtract Packed Integers
func PSUBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Integers
func PSUBWm128int16(X1 []int16, X2 []int16)

// go:noescape

// Unpack High Data
func PUNPCKHBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack High Data
func PUNPCKHDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack High Data
func PUNPCKHQDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack High Data
func PUNPCKHWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack Low Data
func PUNPCKLBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack Low Data
func PUNPCKLDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack Low Data
func PUNPCKLQDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack Low Data
func PUNPCKLWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Logical Exclusive OR
func PXORm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compute Square Roots of Packed Double-Precision Floating-Point Values
func SQRTPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Compute Square Roots of Packed Double-Precision Floating-Point Values
func SQRTPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Compute Square Root of Scalar Double-Precision Floating-Point Value
func SQRTSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Compute Square Root of Scalar Double-Precision Floating-Point Value
func SQRTSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Subtract Packed Double-Precision Floating-Point Values
func SUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Packed Double-Precision Floating-Point Values
func SUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Subtract Scalar Double-Precision Floating-Point Values
func SUBSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Subtract Scalar Double-Precision Floating-Point Values
func SUBSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// Unordered Compare Scalar Double-Precision Floating- Point Values and Set EFLAGS
func UCOMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack and Interleave High Packed Double- Precision Floating-Point Values
func UNPCKHPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func UNPCKLPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func UNPCKLPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// Bitwise Logical XOR for Double-Precision Floating-Point Values
func XORPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// Bitwise Logical XOR for Double-Precision Floating-Point Values
func XORPDm128float64(X1 []float64, X2 []float64)
