package sse2

// go:noescape

// ADDPDm128byte Add Packed Double-Precision Floating-Point Values
func ADDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// ADDPDm128float64 Add Packed Double-Precision Floating-Point Values
func ADDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// ADDSDm64byte Add Scalar Double-Precision Floating-Point Values
func ADDSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// ADDSDm64float64 Add Scalar Double-Precision Floating-Point Values
func ADDSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// ANDNPDm128byte Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func ANDNPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// ANDNPDm128float64 Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func ANDNPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// ANDPDm128byte Bitwise Logical AND of Packed Double Precision Floating-Point Values
func ANDPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// ANDPDm128float64 Bitwise Logical AND of Packed Double Precision Floating-Point Values
func ANDPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// COMISDm64byte Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func COMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// COMISDm64float64 Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func COMISDm64float64(X1 []float64, X2 []float64)

// go:noescape

// DIVPDm128byte Divide Packed Double-Precision Floating-Point Values
func DIVPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// DIVPDm128float64 Divide Packed Double-Precision Floating-Point Values
func DIVPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// DIVSDm64byte Divide Scalar Double-Precision Floating-Point Value
func DIVSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// DIVSDm64float64 Divide Scalar Double-Precision Floating-Point Value
func DIVSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// MASKMOVDQUbyte Store Selected Bytes of Double Quadword
func MASKMOVDQUbyte(X1 []byte, X2 []byte)

// go:noescape

// MAXPDm128byte Maximum of Packed Double-Precision Floating-Point Values
func MAXPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// MAXPDm128float64 Maximum of Packed Double-Precision Floating-Point Values
func MAXPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// MAXSDm64byte Return Maximum Scalar Double-Precision Floating-Point Value
func MAXSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// MAXSDm64float64 Return Maximum Scalar Double-Precision Floating-Point Value
func MAXSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// MINPDm128byte Minimum of Packed Double-Precision Floating-Point Values
func MINPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// MINPDm128float64 Minimum of Packed Double-Precision Floating-Point Values
func MINPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// MINSDm64byte Return Minimum Scalar Double-Precision Floating-Point Value
func MINSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// MINSDm64float64 Return Minimum Scalar Double-Precision Floating-Point Value
func MINSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// MULPDm128byte Multiply Packed Double-Precision Floating-Point Values
func MULPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// MULPDm128float64 Multiply Packed Double-Precision Floating-Point Values
func MULPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// MULSDm64byte Multiply Scalar Double-Precision Floating-Point Value
func MULSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// MULSDm64float64 Multiply Scalar Double-Precision Floating-Point Value
func MULSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// ORPDm128byte Bitwise Logical OR of Packed Double Precision Floating-Point Values
func ORPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// ORPDm128float64 Bitwise Logical OR of Packed Double Precision Floating-Point Values
func ORPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// PACKSSDWm128byte Pack with Signed Saturation
func PACKSSDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PACKSSWBm128byte Pack with Signed Saturation
func PACKSSWBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PACKUSWBm128byte Pack with Unsigned Saturation
func PACKUSWBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDBm128byte Add Packed Integers
func PADDBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDBm128int8 Add Packed Integers
func PADDBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PADDDm128byte Add Packed Integers
func PADDDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDDm128int32 Add Packed Integers
func PADDDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PADDQm128byte Add Packed Integers
func PADDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDQm128int64 Add Packed Integers
func PADDQm128int64(X1 []int64, X2 []int64)

// go:noescape

// PADDSBm128byte Add Packed Signed Integers with Signed Saturation
func PADDSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDSBm128int8 Add Packed Signed Integers with Signed Saturation
func PADDSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PADDSWm128byte Add Packed Signed Integers with Signed Saturation
func PADDSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDSWm128int16 Add Packed Signed Integers with Signed Saturation
func PADDSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PADDUSBm128byte Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDUSBm128uint8 Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// PADDUSWm128byte Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDUSWm128uint16 Add Packed Unsigned Integers with Unsigned Saturation
func PADDUSWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// PADDWm128byte Add Packed Integers
func PADDWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PADDWm128int16 Add Packed Integers
func PADDWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PANDm128byte Logical AND
func PANDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PANDNm128byte Logical AND NOT
func PANDNm128byte(X1 []byte, X2 []byte)

// go:noescape

// PAVGBm128byte Average Packed Integers
func PAVGBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PAVGBm128int8 Average Packed Integers
func PAVGBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PAVGWm128byte Average Packed Integers
func PAVGWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PAVGWm128int16 Average Packed Integers
func PAVGWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PCMPEQBm128byte Compare Packed Data for Equal
func PCMPEQBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPEQDm128byte Compare Packed Data for Equal
func PCMPEQDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPEQWm128byte Compare Packed Data for Equal
func PCMPEQWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPGTBm128byte Compare Packed Signed Integers for Greater Than
func PCMPGTBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPGTBm128int8 Compare Packed Signed Integers for Greater Than
func PCMPGTBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PCMPGTDm128byte Compare Packed Signed Integers for Greater Than
func PCMPGTDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPGTDm128int32 Compare Packed Signed Integers for Greater Than
func PCMPGTDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PCMPGTWm128byte Compare Packed Signed Integers for Greater Than
func PCMPGTWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PCMPGTWm128int16 Compare Packed Signed Integers for Greater Than
func PCMPGTWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PMADDWDm128byte Multiply and Add Packed Integers
func PMADDWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMADDWDm128int32 Multiply and Add Packed Integers
func PMADDWDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PMAXSWm128byte Maximum of Packed Signed Integers
func PMAXSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMAXSWm128int16 Maximum of Packed Signed Integers
func PMAXSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PMAXUBm128byte Maximum of Packed Unsigned Integers
func PMAXUBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMAXUBm128uint8 Maximum of Packed Unsigned Integers
func PMAXUBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// PMINSWm128byte Minimum of Packed Signed Integers
func PMINSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMINSWm128int16 Minimum of Packed Signed Integers
func PMINSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PMINUBm128byte Minimum of Packed Unsigned Integers
func PMINUBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMINUBm128uint8 Minimum of Packed Unsigned Integers
func PMINUBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// PMULHUWm128byte Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMULHUWm128uint16 Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// PMULHWm128byte Multiply Packed Signed Integers and Store High Result
func PMULHWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMULHWm128int16 Multiply Packed Signed Integers and Store High Result
func PMULHWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PMULLWm128byte Multiply Packed Signed Integers and Store Low Result
func PMULLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMULLWm128int16 Multiply Packed Signed Integers and Store Low Result
func PMULLWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PMULUDQm128byte Multiply Packed Unsigned Doubleword Integers
func PMULUDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PMULUDQm128int64 Multiply Packed Unsigned Doubleword Integers
func PMULUDQm128int64(X1 []int64, X2 []int64)

// go:noescape

// PORm128byte Bitwise Logical OR
func PORm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSADBWm128byte Compute Sum of Absolute Differences
func PSADBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSLLDm128byte Shift Packed Data Left Logical
func PSLLDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSLLQm128byte Shift Packed Data Left Logical
func PSLLQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSLLWm128byte Shift Packed Data Left Logical
func PSLLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSRADm128byte Shift Packed Data Right Arithmetic
func PSRADm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSRAWm128byte Shift Packed Data Right Arithmetic
func PSRAWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSRLDm128byte Shift Packed Data Right Logical
func PSRLDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSRLQm128byte Shift Packed Data Right Logical
func PSRLQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSRLWm128byte Shift Packed Data Right Logical
func PSRLWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBBm128byte Subtract Packed Integers
func PSUBBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBBm128int8 Subtract Packed Integers
func PSUBBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PSUBDm128byte Subtract Packed Integers
func PSUBDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBDm128int32 Subtract Packed Integers
func PSUBDm128int32(X1 []int32, X2 []int32)

// go:noescape

// PSUBQm128byte Subtract Packed Quadword Integers
func PSUBQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBQm128int64 Subtract Packed Quadword Integers
func PSUBQm128int64(X1 []int64, X2 []int64)

// go:noescape

// PSUBSBm128byte Subtract Packed Signed Integers with Signed Saturation
func PSUBSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBSBm128int8 Subtract Packed Signed Integers with Signed Saturation
func PSUBSBm128int8(X1 []int8, X2 []int8)

// go:noescape

// PSUBSWm128byte Subtract Packed Signed Integers with Signed Saturation
func PSUBSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBSWm128int16 Subtract Packed Signed Integers with Signed Saturation
func PSUBSWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PSUBUSBm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBUSBm128uint8 Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSBm128uint8(X1 []uint8, X2 []uint8)

// go:noescape

// PSUBUSWm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBUSWm128uint16 Subtract Packed Unsigned Integers with Unsigned Saturation
func PSUBUSWm128uint16(X1 []uint16, X2 []uint16)

// go:noescape

// PSUBWm128byte Subtract Packed Integers
func PSUBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PSUBWm128int16 Subtract Packed Integers
func PSUBWm128int16(X1 []int16, X2 []int16)

// go:noescape

// PUNPCKHBWm128byte Unpack High Data
func PUNPCKHBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKHDQm128byte Unpack High Data
func PUNPCKHDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKHQDQm128byte Unpack High Data
func PUNPCKHQDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKHWDm128byte Unpack High Data
func PUNPCKHWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKLBWm128byte Unpack Low Data
func PUNPCKLBWm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKLDQm128byte Unpack Low Data
func PUNPCKLDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKLQDQm128byte Unpack Low Data
func PUNPCKLQDQm128byte(X1 []byte, X2 []byte)

// go:noescape

// PUNPCKLWDm128byte Unpack Low Data
func PUNPCKLWDm128byte(X1 []byte, X2 []byte)

// go:noescape

// PXORm128byte Logical Exclusive OR
func PXORm128byte(X1 []byte, X2 []byte)

// go:noescape

// SQRTPDm128byte Square Root of Double-Precision Floating-Point Values
func SQRTPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// SQRTPDm128float64 Square Root of Double-Precision Floating-Point Values
func SQRTPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// SQRTSDm64byte Compute Square Root of Scalar Double-Precision Floating-Point Value
func SQRTSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// SQRTSDm64float64 Compute Square Root of Scalar Double-Precision Floating-Point Value
func SQRTSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// SUBPDm128byte Subtract Packed Double-Precision Floating-Point Values
func SUBPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// SUBPDm128float64 Subtract Packed Double-Precision Floating-Point Values
func SUBPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// SUBSDm64byte Subtract Scalar Double-Precision Floating-Point Value
func SUBSDm64byte(X1 []byte, X2 []byte)

// go:noescape

// SUBSDm64float64 Subtract Scalar Double-Precision Floating-Point Value
func SUBSDm64float64(X1 []float64, X2 []float64)

// go:noescape

// UCOMISDm64byte Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func UCOMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// UCOMISDm64float64 Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func UCOMISDm64float64(X1 []float64, X2 []float64)

// go:noescape

// UNPCKHPDm128byte Unpack and Interleave High Packed Double-Precision Floating-Point Values
func UNPCKHPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// UNPCKHPDm128float64 Unpack and Interleave High Packed Double-Precision Floating-Point Values
func UNPCKHPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// UNPCKLPDm128byte Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func UNPCKLPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// UNPCKLPDm128float64 Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func UNPCKLPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// XORPDm128byte Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func XORPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// XORPDm128float64 Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func XORPDm128float64(X1 []float64, X2 []float64)
