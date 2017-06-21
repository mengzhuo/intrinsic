package avx



// go:noescape

// Add Packed Double-Precision Floating-Point Values
func VADDPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Double-Precision Floating-Point Values
func VADDPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Add Packed Single-Precision Floating-Point Values
func VADDPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Single-Precision Floating-Point Values
func VADDPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Add Scalar Double-Precision Floating-Point Values
func VADDSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Scalar Double-Precision Floating-Point Values
func VADDSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Add Scalar Single-Precision Floating-Point Values
func VADDSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Scalar Single-Precision Floating-Point Values
func VADDSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Double-FP Add/Subtract
func VADDSUBPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Double-FP Add/Subtract
func VADDSUBPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Packed Single-FP Add/Subtract
func VADDSUBPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Single-FP Add/Subtract
func VADDSUBPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func VANDNPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func VANDNPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Bitwise Logical AND NOT of Packed Single Precision Floating-Point Values
func VANDNPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical AND NOT of Packed Single Precision Floating-Point Values
func VANDNPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Bitwise Logical AND of Packed Double Precision Floating-Point Values
func VANDPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical AND of Packed Double Precision Floating-Point Values
func VANDPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Bitwise Logical AND of Packed Single Precision Floating-Point Values
func VANDPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical AND of Packed Single Precision Floating-Point Values
func VANDPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func VCOMISDm64byte(X1 []byte,  X2 []byte)

// go:noescape

// Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func VCOMISDm64float64(X1 []float64,  X2 []float64)


// go:noescape

// Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS
func VCOMISSm32byte(X1 []byte,  X2 []byte)

// go:noescape

// Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS
func VCOMISSm32float32(X1 []float32,  X2 []float32)


// go:noescape

// Divide Packed Double-Precision Floating-Point Values
func VDIVPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Divide Packed Double-Precision Floating-Point Values
func VDIVPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Divide Packed Single-Precision Floating-Point Values
func VDIVPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Divide Packed Single-Precision Floating-Point Values
func VDIVPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Divide Scalar Double-Precision Floating-Point Value
func VDIVSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Divide Scalar Double-Precision Floating-Point Value
func VDIVSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Divide Scalar Single-Precision Floating-Point Values
func VDIVSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Divide Scalar Single-Precision Floating-Point Values
func VDIVSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Double-FP Horizontal Add
func VHADDPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Double-FP Horizontal Add
func VHADDPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Packed Single-FP Horizontal Add
func VHADDPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Single-FP Horizontal Add
func VHADDPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Double-FP Horizontal Subtract
func VHSUBPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Double-FP Horizontal Subtract
func VHSUBPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Packed Single-FP Horizontal Subtract
func VHSUBPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Packed Single-FP Horizontal Subtract
func VHSUBPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Store Selected Bytes of Double Quadword
func VMASKMOVDQUbyte(X1 []byte,  X2 []byte)


// go:noescape

// Maximum of Packed Double-Precision Floating-Point Values
func VMAXPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Double-Precision Floating-Point Values
func VMAXPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Maximum of Packed Single-Precision Floating-Point Values
func VMAXPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Single-Precision Floating-Point Values
func VMAXPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Return Maximum Scalar Double-Precision Floating-Point Value
func VMAXSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Return Maximum Scalar Double-Precision Floating-Point Value
func VMAXSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Return Maximum Scalar Single-Precision Floating-Point Value
func VMAXSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Return Maximum Scalar Single-Precision Floating-Point Value
func VMAXSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Minimum of Packed Double-Precision Floating-Point Values
func VMINPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Double-Precision Floating-Point Values
func VMINPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Minimum of Packed Single-Precision Floating-Point Values
func VMINPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Single-Precision Floating-Point Values
func VMINPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Return Minimum Scalar Double-Precision Floating-Point Value
func VMINSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Return Minimum Scalar Double-Precision Floating-Point Value
func VMINSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Return Minimum Scalar Single-Precision Floating-Point Value
func VMINSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Return Minimum Scalar Single-Precision Floating-Point Value
func VMINSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Multiply Packed Double-Precision Floating-Point Values
func VMULPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Double-Precision Floating-Point Values
func VMULPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Multiply Packed Single-Precision Floating-Point Values
func VMULPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Single-Precision Floating-Point Values
func VMULPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Multiply Scalar Double-Precision Floating-Point Value
func VMULSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Scalar Double-Precision Floating-Point Value
func VMULSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Multiply Scalar Single-Precision Floating-Point Values
func VMULSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Scalar Single-Precision Floating-Point Values
func VMULSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Bitwise Logical OR of Packed Double Precision Floating-Point Values
func VORPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical OR of Packed Double Precision Floating-Point Values
func VORPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Bitwise Logical OR of Packed Single Precision Floating-Point Values
func VORPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical OR of Packed Single Precision Floating-Point Values
func VORPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Absolute Value
func VPABSBm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Absolute Value
func VPABSDm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Absolute Value
func VPABSWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Pack with Signed Saturation
func VPACKSSDWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Pack with Signed Saturation
func VPACKSSWBm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Pack with Unsigned Saturation
func VPACKUSDWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Pack with Unsigned Saturation
func VPACKUSWBm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Add Packed Integers
func VPADDBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Integers
func VPADDBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Add Packed Integers
func VPADDDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Integers
func VPADDDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Add Packed Integers
func VPADDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Integers
func VPADDQm128int64(X1 []int64,  X2 []int64,  X3 []int64)


// go:noescape

// Add Packed Signed Integers with Signed Saturation
func VPADDSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func VPADDSBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Add Packed Signed Integers with Signed Saturation
func VPADDSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Signed Integers with Signed Saturation
func VPADDSWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSBm128uint8(X1 []uint8,  X2 []uint8,  X3 []uint8)


// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSWm128uint16(X1 []uint16,  X2 []uint16,  X3 []uint16)


// go:noescape

// Add Packed Integers
func VPADDWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Add Packed Integers
func VPADDWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Logical AND
func VPANDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Logical AND NOT
func VPANDNm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Average Packed Integers
func VPAVGBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Average Packed Integers
func VPAVGBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Average Packed Integers
func VPAVGWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Average Packed Integers
func VPAVGWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Compare Packed Data for Equal
func VPCMPEQBm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compare Packed Data for Equal
func VPCMPEQDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compare Packed Qword Data for Equal
func VPCMPEQQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compare Packed Data for Equal
func VPCMPEQWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Compare Packed Data for Greater Than
func VPCMPGTQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compare Packed Signed Integers for Greater Than
func VPCMPGTWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Permute In-Lane of Pairs of Double-Precision Floating-Point Values
func VPERMILPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Permute In-Lane of Pairs of Double-Precision Floating-Point Values
func VPERMILPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Permute In-Lane of Quadruples of Single-Precision Floating-Point Values
func VPERMILPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Permute In-Lane of Quadruples of Single-Precision Floating-Point Values
func VPERMILPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Horizontal Add
func VPHADDDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Horizontal Add and Saturate
func VPHADDSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Horizontal Add
func VPHADDWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Horizontal Word Minimum
func VPHMINPOSUWm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Horizontal Subtract
func VPHSUBDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Horizontal Subtract and Saturate
func VPHSUBSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Horizontal Subtract
func VPHSUBWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Multiply and Add Packed Signed and Unsigned Bytes
func VPMADDUBSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Multiply and Add Packed Integers
func VPMADDWDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply and Add Packed Integers
func VPMADDWDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Signed Integers
func VPMAXSWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUBm128uint8(X1 []uint8,  X2 []uint8,  X3 []uint8)


// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUDm128uint32(X1 []uint32,  X2 []uint32,  X3 []uint32)


// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Maximum of Packed Unsigned Integers
func VPMAXUWm128uint16(X1 []uint16,  X2 []uint16,  X3 []uint16)


// go:noescape

// Minimum of Packed Signed Integers
func VPMINSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Signed Integers
func VPMINSBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Minimum of Packed Signed Integers
func VPMINSDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Signed Integers
func VPMINSDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Minimum of Packed Signed Integers
func VPMINSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Signed Integers
func VPMINSWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUBm128uint8(X1 []uint8,  X2 []uint8,  X3 []uint8)


// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUDm128uint32(X1 []uint32,  X2 []uint32,  X3 []uint32)


// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Minimum of Packed Unsigned Integers
func VPMINUWm128uint16(X1 []uint16,  X2 []uint16,  X3 []uint16)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXBDm32byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXBQm16byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXBWm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXDQm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXWDm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Sign Extend
func VPMOVSXWQm32byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXBDm32byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXBQm16byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXBWm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXDQm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXWDm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Move with Zero Extend
func VPMOVZXWQm32byte(X1 []byte,  X2 []byte)


// go:noescape

// Multiply Packed Doubleword Integers
func VPMULDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Doubleword Integers
func VPMULDQm128int64(X1 []int64,  X2 []int64,  X3 []int64)


// go:noescape

// Packed Multiply High with Round and Scale
func VPMULHRSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Multiply Packed Unsigned Integers and Store High Result
func VPMULHUWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Unsigned Integers and Store High Result
func VPMULHUWm128uint16(X1 []uint16,  X2 []uint16,  X3 []uint16)


// go:noescape

// Multiply Packed Signed Integers and Store High Result
func VPMULHWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Signed Integers and Store High Result
func VPMULHWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Multiply Packed Integers and Store Low Result
func VPMULLDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Integers and Store Low Result
func VPMULLDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Multiply Packed Signed Integers and Store Low Result
func VPMULLWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Signed Integers and Store Low Result
func VPMULLWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Multiply Packed Unsigned Doubleword Integers
func VPMULUDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Multiply Packed Unsigned Doubleword Integers
func VPMULUDQm128int64(X1 []int64,  X2 []int64,  X3 []int64)


// go:noescape

// Bitwise Logical OR
func VPORm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compute Sum of Absolute Differences
func VPSADBWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed Shuffle Bytes
func VPSHUFBm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed SIGN
func VPSIGNBm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed SIGN
func VPSIGNDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Packed SIGN
func VPSIGNWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Left Logical
func VPSLLDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Left Logical
func VPSLLQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Left Logical
func VPSLLWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Right Arithmetic
func VPSRADm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Right Arithmetic
func VPSRAWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Right Logical
func VPSRLDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Right Logical
func VPSRLQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Shift Packed Data Right Logical
func VPSRLWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Subtract Packed Integers
func VPSUBBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Integers
func VPSUBBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Subtract Packed Integers
func VPSUBDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Integers
func VPSUBDm128int32(X1 []int32,  X2 []int32,  X3 []int32)


// go:noescape

// Subtract Packed Quadword Integers
func VPSUBQm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Quadword Integers
func VPSUBQm128int64(X1 []int64,  X2 []int64,  X3 []int64)


// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func VPSUBSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func VPSUBSBm128int8(X1 []int8,  X2 []int8,  X3 []int8)


// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func VPSUBSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Signed Integers with Signed Saturation
func VPSUBSWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSBm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSBm128uint8(X1 []uint8,  X2 []uint8,  X3 []uint8)


// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSWm128uint16(X1 []uint16,  X2 []uint16,  X3 []uint16)


// go:noescape

// Subtract Packed Integers
func VPSUBWm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Integers
func VPSUBWm128int16(X1 []int16,  X2 []int16,  X3 []int16)


// go:noescape

// PTEST- Logical Compare
func VPTESTm128byte(X1 []byte,  X2 []byte)


// go:noescape

// PTEST- Logical Compare
func VPTESTm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Unpack High Data
func VPUNPCKHBWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack High Data
func VPUNPCKHDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack High Data
func VPUNPCKHQDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack High Data
func VPUNPCKHWDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack Low Data
func VPUNPCKLBWm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack Low Data
func VPUNPCKLDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack Low Data
func VPUNPCKLQDQm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Unpack Low Data
func VPUNPCKLWDm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Logical Exclusive OR
func VPXORm128byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Compute Reciprocal of Scalar Single-Precision Floating-Point Values
func VRCPSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compute Reciprocal of Scalar Single-Precision Floating-Point Values
func VRCPSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value
func VRSQRTSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value
func VRSQRTSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Square Root of Double-Precision Floating-Point Values
func VSQRTPDm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Square Root of Double-Precision Floating-Point Values
func VSQRTPDm128float64(X1 []float64,  X2 []float64)


// go:noescape

// Square Root of Double-Precision Floating-Point Values
func VSQRTPDm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Square Root of Double-Precision Floating-Point Values
func VSQRTPDm256float64(Y1 []float64,  Y2 []float64)


// go:noescape

// Square Root of Single-Precision Floating-Point Values
func VSQRTPSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Square Root of Single-Precision Floating-Point Values
func VSQRTPSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Square Root of Single-Precision Floating-Point Values
func VSQRTPSm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Square Root of Single-Precision Floating-Point Values
func VSQRTPSm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Compute Square Root of Scalar Double-Precision Floating-Point Value
func VSQRTSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Compute Square Root of Scalar Double-Precision Floating-Point Value
func VSQRTSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Compute Square Root of Scalar Single-Precision Value
func VSQRTSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)


// go:noescape

// Subtract Packed Double-Precision Floating-Point Values
func VSUBPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Double-Precision Floating-Point Values
func VSUBPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Subtract Packed Single-Precision Floating-Point Values
func VSUBPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Packed Single-Precision Floating-Point Values
func VSUBPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Subtract Scalar Double-Precision Floating-Point Value
func VSUBSDm64byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Scalar Double-Precision Floating-Point Value
func VSUBSDm64float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Subtract Scalar Single-Precision Floating-Point Value
func VSUBSSm32byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Subtract Scalar Single-Precision Floating-Point Value
func VSUBSSm32float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Packed Bit Test
func VTESTPDm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Bit Test
func VTESTPDm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Packed Bit Test
func VTESTPSm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Packed Bit Test
func VTESTPSm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func VUCOMISDm64byte(X1 []byte,  X2 []byte)

// go:noescape

// Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func VUCOMISDm64float64(X1 []float64,  X2 []float64)


// go:noescape

// Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS
func VUCOMISSm32byte(X1 []byte,  X2 []byte)

// go:noescape

// Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS
func VUCOMISSm32float32(X1 []float32,  X2 []float32)


// go:noescape

// Unpack and Interleave High Packed Double-Precision Floating-Point Values
func VUNPCKHPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Unpack and Interleave High Packed Double-Precision Floating-Point Values
func VUNPCKHPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Unpack and Interleave High Packed Single-Precision Floating-Point Values
func VUNPCKHPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Unpack and Interleave High Packed Single-Precision Floating-Point Values
func VUNPCKHPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func VUNPCKLPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func VUNPCKLPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Unpack and Interleave Low Packed Single-Precision Floating-Point Values
func VUNPCKLPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Unpack and Interleave Low Packed Single-Precision Floating-Point Values
func VUNPCKLPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)


// go:noescape

// Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func VXORPDm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func VXORPDm128float64(X1 []float64,  X2 []float64,  X3 []float64)


// go:noescape

// Bitwise Logical XOR of Packed Single Precision Floating-Point Values
func VXORPSm128byte(X1 []byte,  X2 []byte,  X3 []byte)

// go:noescape

// Bitwise Logical XOR of Packed Single Precision Floating-Point Values
func VXORPSm128float32(X1 []float32,  X2 []float32,  X3 []float32)

