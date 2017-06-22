package avx

// go:noescape

// VADDPDm128byte Add Packed Double-Precision Floating-Point Values
func VADDPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDPDm128float64 Add Packed Double-Precision Floating-Point Values
func VADDPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VADDPSm128byte Add Packed Single-Precision Floating-Point Values
func VADDPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDPSm128float32 Add Packed Single-Precision Floating-Point Values
func VADDPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VADDSDm64byte Add Scalar Double-Precision Floating-Point Values
func VADDSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDSDm64float64 Add Scalar Double-Precision Floating-Point Values
func VADDSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VADDSSm32byte Add Scalar Single-Precision Floating-Point Values
func VADDSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDSSm32float32 Add Scalar Single-Precision Floating-Point Values
func VADDSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VADDSUBPDm128byte Packed Double-FP Add/Subtract
func VADDSUBPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDSUBPDm128float64 Packed Double-FP Add/Subtract
func VADDSUBPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VADDSUBPSm128byte Packed Single-FP Add/Subtract
func VADDSUBPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VADDSUBPSm128float32 Packed Single-FP Add/Subtract
func VADDSUBPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VANDNPDm128byte Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func VANDNPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VANDNPDm128float64 Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
func VANDNPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VANDNPSm128byte Bitwise Logical AND NOT of Packed Single Precision Floating-Point Values
func VANDNPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VANDNPSm128float32 Bitwise Logical AND NOT of Packed Single Precision Floating-Point Values
func VANDNPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VANDPDm128byte Bitwise Logical AND of Packed Double Precision Floating-Point Values
func VANDPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VANDPDm128float64 Bitwise Logical AND of Packed Double Precision Floating-Point Values
func VANDPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VANDPSm128byte Bitwise Logical AND of Packed Single Precision Floating-Point Values
func VANDPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VANDPSm128float32 Bitwise Logical AND of Packed Single Precision Floating-Point Values
func VANDPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VCOMISDm64byte Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func VCOMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// VCOMISDm64float64 Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
func VCOMISDm64float64(X1 []float64, X2 []float64)

// go:noescape

// VCOMISSm32byte Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS
func VCOMISSm32byte(X1 []byte, X2 []byte)

// go:noescape

// VCOMISSm32float32 Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS
func VCOMISSm32float32(X1 []float32, X2 []float32)

// go:noescape

// VDIVPDm128byte Divide Packed Double-Precision Floating-Point Values
func VDIVPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VDIVPDm128float64 Divide Packed Double-Precision Floating-Point Values
func VDIVPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VDIVPSm128byte Divide Packed Single-Precision Floating-Point Values
func VDIVPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VDIVPSm128float32 Divide Packed Single-Precision Floating-Point Values
func VDIVPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VDIVSDm64byte Divide Scalar Double-Precision Floating-Point Value
func VDIVSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VDIVSDm64float64 Divide Scalar Double-Precision Floating-Point Value
func VDIVSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VDIVSSm32byte Divide Scalar Single-Precision Floating-Point Values
func VDIVSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VDIVSSm32float32 Divide Scalar Single-Precision Floating-Point Values
func VDIVSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VHADDPDm128byte Packed Double-FP Horizontal Add
func VHADDPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VHADDPDm128float64 Packed Double-FP Horizontal Add
func VHADDPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VHADDPSm128byte Packed Single-FP Horizontal Add
func VHADDPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VHADDPSm128float32 Packed Single-FP Horizontal Add
func VHADDPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VHSUBPDm128byte Packed Double-FP Horizontal Subtract
func VHSUBPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VHSUBPDm128float64 Packed Double-FP Horizontal Subtract
func VHSUBPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VHSUBPSm128byte Packed Single-FP Horizontal Subtract
func VHSUBPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VHSUBPSm128float32 Packed Single-FP Horizontal Subtract
func VHSUBPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMASKMOVDQUbyte Store Selected Bytes of Double Quadword
func VMASKMOVDQUbyte(X1 []byte, X2 []byte)

// go:noescape

// VMAXPDm128byte Maximum of Packed Double-Precision Floating-Point Values
func VMAXPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMAXPDm128float64 Maximum of Packed Double-Precision Floating-Point Values
func VMAXPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMAXPSm128byte Maximum of Packed Single-Precision Floating-Point Values
func VMAXPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMAXPSm128float32 Maximum of Packed Single-Precision Floating-Point Values
func VMAXPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMAXSDm64byte Return Maximum Scalar Double-Precision Floating-Point Value
func VMAXSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMAXSDm64float64 Return Maximum Scalar Double-Precision Floating-Point Value
func VMAXSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMAXSSm32byte Return Maximum Scalar Single-Precision Floating-Point Value
func VMAXSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMAXSSm32float32 Return Maximum Scalar Single-Precision Floating-Point Value
func VMAXSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMINPDm128byte Minimum of Packed Double-Precision Floating-Point Values
func VMINPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMINPDm128float64 Minimum of Packed Double-Precision Floating-Point Values
func VMINPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMINPSm128byte Minimum of Packed Single-Precision Floating-Point Values
func VMINPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMINPSm128float32 Minimum of Packed Single-Precision Floating-Point Values
func VMINPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMINSDm64byte Return Minimum Scalar Double-Precision Floating-Point Value
func VMINSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMINSDm64float64 Return Minimum Scalar Double-Precision Floating-Point Value
func VMINSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMINSSm32byte Return Minimum Scalar Single-Precision Floating-Point Value
func VMINSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMINSSm32float32 Return Minimum Scalar Single-Precision Floating-Point Value
func VMINSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMULPDm128byte Multiply Packed Double-Precision Floating-Point Values
func VMULPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMULPDm128float64 Multiply Packed Double-Precision Floating-Point Values
func VMULPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMULPSm128byte Multiply Packed Single-Precision Floating-Point Values
func VMULPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMULPSm128float32 Multiply Packed Single-Precision Floating-Point Values
func VMULPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VMULSDm64byte Multiply Scalar Double-Precision Floating-Point Value
func VMULSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMULSDm64float64 Multiply Scalar Double-Precision Floating-Point Value
func VMULSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VMULSSm32byte Multiply Scalar Single-Precision Floating-Point Values
func VMULSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VMULSSm32float32 Multiply Scalar Single-Precision Floating-Point Values
func VMULSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VORPDm128byte Bitwise Logical OR of Packed Double Precision Floating-Point Values
func VORPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VORPDm128float64 Bitwise Logical OR of Packed Double Precision Floating-Point Values
func VORPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VORPSm128byte Bitwise Logical OR of Packed Single Precision Floating-Point Values
func VORPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VORPSm128float32 Bitwise Logical OR of Packed Single Precision Floating-Point Values
func VORPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VPABSBm128byte Packed Absolute Value
func VPABSBm128byte(X1 []byte, X2 []byte)

// go:noescape

// VPABSDm128byte Packed Absolute Value
func VPABSDm128byte(X1 []byte, X2 []byte)

// go:noescape

// VPABSWm128byte Packed Absolute Value
func VPABSWm128byte(X1 []byte, X2 []byte)

// go:noescape

// VPACKSSDWm128byte Pack with Signed Saturation
func VPACKSSDWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPACKSSWBm128byte Pack with Signed Saturation
func VPACKSSWBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPACKUSDWm128byte Pack with Unsigned Saturation
func VPACKUSDWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPACKUSWBm128byte Pack with Unsigned Saturation
func VPACKUSWBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDBm128byte Add Packed Integers
func VPADDBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDBm128int8 Add Packed Integers
func VPADDBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPADDDm128byte Add Packed Integers
func VPADDDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDDm128int32 Add Packed Integers
func VPADDDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPADDQm128byte Add Packed Integers
func VPADDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDQm128int64 Add Packed Integers
func VPADDQm128int64(X1 []int64, X2 []int64, X3 []int64)

// go:noescape

// VPADDSBm128byte Add Packed Signed Integers with Signed Saturation
func VPADDSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDSBm128int8 Add Packed Signed Integers with Signed Saturation
func VPADDSBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPADDSWm128byte Add Packed Signed Integers with Signed Saturation
func VPADDSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDSWm128int16 Add Packed Signed Integers with Signed Saturation
func VPADDSWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPADDUSBm128byte Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDUSBm128uint8 Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSBm128uint8(X1 []uint8, X2 []uint8, X3 []uint8)

// go:noescape

// VPADDUSWm128byte Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDUSWm128uint16 Add Packed Unsigned Integers with Unsigned Saturation
func VPADDUSWm128uint16(X1 []uint16, X2 []uint16, X3 []uint16)

// go:noescape

// VPADDWm128byte Add Packed Integers
func VPADDWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPADDWm128int16 Add Packed Integers
func VPADDWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPANDm128byte Logical AND
func VPANDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPANDNm128byte Logical AND NOT
func VPANDNm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPAVGBm128byte Average Packed Integers
func VPAVGBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPAVGBm128int8 Average Packed Integers
func VPAVGBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPAVGWm128byte Average Packed Integers
func VPAVGWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPAVGWm128int16 Average Packed Integers
func VPAVGWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPCMPEQBm128byte Compare Packed Data for Equal
func VPCMPEQBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPEQDm128byte Compare Packed Data for Equal
func VPCMPEQDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPEQQm128byte Compare Packed Qword Data for Equal
func VPCMPEQQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPEQWm128byte Compare Packed Data for Equal
func VPCMPEQWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPGTBm128byte Compare Packed Signed Integers for Greater Than
func VPCMPGTBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPGTBm128int8 Compare Packed Signed Integers for Greater Than
func VPCMPGTBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPCMPGTDm128byte Compare Packed Signed Integers for Greater Than
func VPCMPGTDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPGTDm128int32 Compare Packed Signed Integers for Greater Than
func VPCMPGTDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPCMPGTQm128byte Compare Packed Data for Greater Than
func VPCMPGTQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPGTWm128byte Compare Packed Signed Integers for Greater Than
func VPCMPGTWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPCMPGTWm128int16 Compare Packed Signed Integers for Greater Than
func VPCMPGTWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPERMILPDm128byte Permute In-Lane of Pairs of Double-Precision Floating-Point Values
func VPERMILPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPERMILPDm128float64 Permute In-Lane of Pairs of Double-Precision Floating-Point Values
func VPERMILPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VPERMILPSm128byte Permute In-Lane of Quadruples of Single-Precision Floating-Point Values
func VPERMILPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPERMILPSm128float32 Permute In-Lane of Quadruples of Single-Precision Floating-Point Values
func VPERMILPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VPHADDDm128byte Packed Horizontal Add
func VPHADDDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPHADDSWm128byte Packed Horizontal Add and Saturate
func VPHADDSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPHADDWm128byte Packed Horizontal Add
func VPHADDWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPHMINPOSUWm128byte Packed Horizontal Word Minimum
func VPHMINPOSUWm128byte(X1 []byte, X2 []byte)

// go:noescape

// VPHSUBDm128byte Packed Horizontal Subtract
func VPHSUBDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPHSUBSWm128byte Packed Horizontal Subtract and Saturate
func VPHSUBSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPHSUBWm128byte Packed Horizontal Subtract
func VPHSUBWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMADDUBSWm128byte Multiply and Add Packed Signed and Unsigned Bytes
func VPMADDUBSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMADDWDm128byte Multiply and Add Packed Integers
func VPMADDWDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMADDWDm128int32 Multiply and Add Packed Integers
func VPMADDWDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPMAXSBm128byte Maximum of Packed Signed Integers
func VPMAXSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXSBm128int8 Maximum of Packed Signed Integers
func VPMAXSBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPMAXSDm128byte Maximum of Packed Signed Integers
func VPMAXSDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXSDm128int32 Maximum of Packed Signed Integers
func VPMAXSDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPMAXSWm128byte Maximum of Packed Signed Integers
func VPMAXSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXSWm128int16 Maximum of Packed Signed Integers
func VPMAXSWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPMAXUBm128byte Maximum of Packed Unsigned Integers
func VPMAXUBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXUBm128uint8 Maximum of Packed Unsigned Integers
func VPMAXUBm128uint8(X1 []uint8, X2 []uint8, X3 []uint8)

// go:noescape

// VPMAXUDm128byte Maximum of Packed Unsigned Integers
func VPMAXUDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXUDm128uint32 Maximum of Packed Unsigned Integers
func VPMAXUDm128uint32(X1 []uint32, X2 []uint32, X3 []uint32)

// go:noescape

// VPMAXUWm128byte Maximum of Packed Unsigned Integers
func VPMAXUWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMAXUWm128uint16 Maximum of Packed Unsigned Integers
func VPMAXUWm128uint16(X1 []uint16, X2 []uint16, X3 []uint16)

// go:noescape

// VPMINSBm128byte Minimum of Packed Signed Integers
func VPMINSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINSBm128int8 Minimum of Packed Signed Integers
func VPMINSBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPMINSDm128byte Minimum of Packed Signed Integers
func VPMINSDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINSDm128int32 Minimum of Packed Signed Integers
func VPMINSDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPMINSWm128byte Minimum of Packed Signed Integers
func VPMINSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINSWm128int16 Minimum of Packed Signed Integers
func VPMINSWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPMINUBm128byte Minimum of Packed Unsigned Integers
func VPMINUBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINUBm128uint8 Minimum of Packed Unsigned Integers
func VPMINUBm128uint8(X1 []uint8, X2 []uint8, X3 []uint8)

// go:noescape

// VPMINUDm128byte Minimum of Packed Unsigned Integers
func VPMINUDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINUDm128uint32 Minimum of Packed Unsigned Integers
func VPMINUDm128uint32(X1 []uint32, X2 []uint32, X3 []uint32)

// go:noescape

// VPMINUWm128byte Minimum of Packed Unsigned Integers
func VPMINUWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMINUWm128uint16 Minimum of Packed Unsigned Integers
func VPMINUWm128uint16(X1 []uint16, X2 []uint16, X3 []uint16)

// go:noescape

// VPMOVSXBDm32byte Packed Move with Sign Extend
func VPMOVSXBDm32byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVSXBQm16byte Packed Move with Sign Extend
func VPMOVSXBQm16byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVSXBWm64byte Packed Move with Sign Extend
func VPMOVSXBWm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVSXDQm64byte Packed Move with Sign Extend
func VPMOVSXDQm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVSXWDm64byte Packed Move with Sign Extend
func VPMOVSXWDm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVSXWQm32byte Packed Move with Sign Extend
func VPMOVSXWQm32byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXBDm32byte Packed Move with Zero Extend
func VPMOVZXBDm32byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXBQm16byte Packed Move with Zero Extend
func VPMOVZXBQm16byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXBWm64byte Packed Move with Zero Extend
func VPMOVZXBWm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXDQm64byte Packed Move with Zero Extend
func VPMOVZXDQm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXWDm64byte Packed Move with Zero Extend
func VPMOVZXWDm64byte(X1 []byte, X2 []byte)

// go:noescape

// VPMOVZXWQm32byte Packed Move with Zero Extend
func VPMOVZXWQm32byte(X1 []byte, X2 []byte)

// go:noescape

// VPMULDQm128byte Multiply Packed Doubleword Integers
func VPMULDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULDQm128int64 Multiply Packed Doubleword Integers
func VPMULDQm128int64(X1 []int64, X2 []int64, X3 []int64)

// go:noescape

// VPMULHRSWm128byte Packed Multiply High with Round and Scale
func VPMULHRSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULHUWm128byte Multiply Packed Unsigned Integers and Store High Result
func VPMULHUWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULHUWm128uint16 Multiply Packed Unsigned Integers and Store High Result
func VPMULHUWm128uint16(X1 []uint16, X2 []uint16, X3 []uint16)

// go:noescape

// VPMULHWm128byte Multiply Packed Signed Integers and Store High Result
func VPMULHWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULHWm128int16 Multiply Packed Signed Integers and Store High Result
func VPMULHWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPMULLDm128byte Multiply Packed Integers and Store Low Result
func VPMULLDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULLDm128int32 Multiply Packed Integers and Store Low Result
func VPMULLDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPMULLWm128byte Multiply Packed Signed Integers and Store Low Result
func VPMULLWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULLWm128int16 Multiply Packed Signed Integers and Store Low Result
func VPMULLWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPMULUDQm128byte Multiply Packed Unsigned Doubleword Integers
func VPMULUDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPMULUDQm128int64 Multiply Packed Unsigned Doubleword Integers
func VPMULUDQm128int64(X1 []int64, X2 []int64, X3 []int64)

// go:noescape

// VPORm128byte Bitwise Logical OR
func VPORm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSADBWm128byte Compute Sum of Absolute Differences
func VPSADBWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSHUFBm128byte Packed Shuffle Bytes
func VPSHUFBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSIGNBm128byte Packed SIGN
func VPSIGNBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSIGNDm128byte Packed SIGN
func VPSIGNDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSIGNWm128byte Packed SIGN
func VPSIGNWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSLLDm128byte Shift Packed Data Left Logical
func VPSLLDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSLLQm128byte Shift Packed Data Left Logical
func VPSLLQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSLLWm128byte Shift Packed Data Left Logical
func VPSLLWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSRADm128byte Shift Packed Data Right Arithmetic
func VPSRADm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSRAWm128byte Shift Packed Data Right Arithmetic
func VPSRAWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSRLDm128byte Shift Packed Data Right Logical
func VPSRLDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSRLQm128byte Shift Packed Data Right Logical
func VPSRLQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSRLWm128byte Shift Packed Data Right Logical
func VPSRLWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBBm128byte Subtract Packed Integers
func VPSUBBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBBm128int8 Subtract Packed Integers
func VPSUBBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPSUBDm128byte Subtract Packed Integers
func VPSUBDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBDm128int32 Subtract Packed Integers
func VPSUBDm128int32(X1 []int32, X2 []int32, X3 []int32)

// go:noescape

// VPSUBQm128byte Subtract Packed Quadword Integers
func VPSUBQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBQm128int64 Subtract Packed Quadword Integers
func VPSUBQm128int64(X1 []int64, X2 []int64, X3 []int64)

// go:noescape

// VPSUBSBm128byte Subtract Packed Signed Integers with Signed Saturation
func VPSUBSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBSBm128int8 Subtract Packed Signed Integers with Signed Saturation
func VPSUBSBm128int8(X1 []int8, X2 []int8, X3 []int8)

// go:noescape

// VPSUBSWm128byte Subtract Packed Signed Integers with Signed Saturation
func VPSUBSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBSWm128int16 Subtract Packed Signed Integers with Signed Saturation
func VPSUBSWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPSUBUSBm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSBm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBUSBm128uint8 Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSBm128uint8(X1 []uint8, X2 []uint8, X3 []uint8)

// go:noescape

// VPSUBUSWm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBUSWm128uint16 Subtract Packed Unsigned Integers with Unsigned Saturation
func VPSUBUSWm128uint16(X1 []uint16, X2 []uint16, X3 []uint16)

// go:noescape

// VPSUBWm128byte Subtract Packed Integers
func VPSUBWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPSUBWm128int16 Subtract Packed Integers
func VPSUBWm128int16(X1 []int16, X2 []int16, X3 []int16)

// go:noescape

// VPTESTm128byte PTEST- Logical Compare
func VPTESTm128byte(X1 []byte, X2 []byte)

// go:noescape

// VPTESTm256byte PTEST- Logical Compare
func VPTESTm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VPUNPCKHBWm128byte Unpack High Data
func VPUNPCKHBWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKHDQm128byte Unpack High Data
func VPUNPCKHDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKHQDQm128byte Unpack High Data
func VPUNPCKHQDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKHWDm128byte Unpack High Data
func VPUNPCKHWDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKLBWm128byte Unpack Low Data
func VPUNPCKLBWm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKLDQm128byte Unpack Low Data
func VPUNPCKLDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKLQDQm128byte Unpack Low Data
func VPUNPCKLQDQm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPUNPCKLWDm128byte Unpack Low Data
func VPUNPCKLWDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VPXORm128byte Logical Exclusive OR
func VPXORm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VRCPPSm128byte Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// VRCPPSm128float32 Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// VRCPPSm256byte Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VRCPPSm256float32 Compute Reciprocals of Packed Single-Precision Floating-Point Values
func VRCPPSm256float32(Y1 []float32, Y2 []float32)

// go:noescape

// VRCPSSm32byte Compute Reciprocal of Scalar Single-Precision Floating-Point Values
func VRCPSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VRCPSSm32float32 Compute Reciprocal of Scalar Single-Precision Floating-Point Values
func VRCPSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VRSQRTPSm128byte Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// VRSQRTPSm128float32 Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// VRSQRTPSm256byte Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VRSQRTPSm256float32 Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values
func VRSQRTPSm256float32(Y1 []float32, Y2 []float32)

// go:noescape

// VRSQRTSSm32byte Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value
func VRSQRTSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VRSQRTSSm32float32 Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value
func VRSQRTSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VSQRTPDm128byte Square Root of Double-Precision Floating-Point Values
func VSQRTPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// VSQRTPDm128float64 Square Root of Double-Precision Floating-Point Values
func VSQRTPDm128float64(X1 []float64, X2 []float64)

// go:noescape

// VSQRTPDm256byte Square Root of Double-Precision Floating-Point Values
func VSQRTPDm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VSQRTPDm256float64 Square Root of Double-Precision Floating-Point Values
func VSQRTPDm256float64(Y1 []float64, Y2 []float64)

// go:noescape

// VSQRTPSm128byte Square Root of Single-Precision Floating-Point Values
func VSQRTPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// VSQRTPSm128float32 Square Root of Single-Precision Floating-Point Values
func VSQRTPSm128float32(X1 []float32, X2 []float32)

// go:noescape

// VSQRTPSm256byte Square Root of Single-Precision Floating-Point Values
func VSQRTPSm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VSQRTPSm256float32 Square Root of Single-Precision Floating-Point Values
func VSQRTPSm256float32(Y1 []float32, Y2 []float32)

// go:noescape

// VSQRTSDm64byte Compute Square Root of Scalar Double-Precision Floating-Point Value
func VSQRTSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSQRTSDm64float64 Compute Square Root of Scalar Double-Precision Floating-Point Value
func VSQRTSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VSQRTSSm32byte Compute Square Root of Scalar Single-Precision Value
func VSQRTSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSUBPDm128byte Subtract Packed Double-Precision Floating-Point Values
func VSUBPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSUBPDm128float64 Subtract Packed Double-Precision Floating-Point Values
func VSUBPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VSUBPSm128byte Subtract Packed Single-Precision Floating-Point Values
func VSUBPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSUBPSm128float32 Subtract Packed Single-Precision Floating-Point Values
func VSUBPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VSUBSDm64byte Subtract Scalar Double-Precision Floating-Point Value
func VSUBSDm64byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSUBSDm64float64 Subtract Scalar Double-Precision Floating-Point Value
func VSUBSDm64float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VSUBSSm32byte Subtract Scalar Single-Precision Floating-Point Value
func VSUBSSm32byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VSUBSSm32float32 Subtract Scalar Single-Precision Floating-Point Value
func VSUBSSm32float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VTESTPDm128byte Packed Bit Test
func VTESTPDm128byte(X1 []byte, X2 []byte)

// go:noescape

// VTESTPDm256byte Packed Bit Test
func VTESTPDm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VTESTPSm256byte Packed Bit Test
func VTESTPSm256byte(Y1 []byte, Y2 []byte)

// go:noescape

// VTESTPSm128byte Packed Bit Test
func VTESTPSm128byte(X1 []byte, X2 []byte)

// go:noescape

// VUCOMISDm64byte Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func VUCOMISDm64byte(X1 []byte, X2 []byte)

// go:noescape

// VUCOMISDm64float64 Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
func VUCOMISDm64float64(X1 []float64, X2 []float64)

// go:noescape

// VUCOMISSm32byte Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS
func VUCOMISSm32byte(X1 []byte, X2 []byte)

// go:noescape

// VUCOMISSm32float32 Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS
func VUCOMISSm32float32(X1 []float32, X2 []float32)

// go:noescape

// VUNPCKHPDm128byte Unpack and Interleave High Packed Double-Precision Floating-Point Values
func VUNPCKHPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VUNPCKHPDm128float64 Unpack and Interleave High Packed Double-Precision Floating-Point Values
func VUNPCKHPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VUNPCKHPSm128byte Unpack and Interleave High Packed Single-Precision Floating-Point Values
func VUNPCKHPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VUNPCKHPSm128float32 Unpack and Interleave High Packed Single-Precision Floating-Point Values
func VUNPCKHPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VUNPCKLPDm128byte Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func VUNPCKLPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VUNPCKLPDm128float64 Unpack and Interleave Low Packed Double-Precision Floating-Point Values
func VUNPCKLPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VUNPCKLPSm128byte Unpack and Interleave Low Packed Single-Precision Floating-Point Values
func VUNPCKLPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VUNPCKLPSm128float32 Unpack and Interleave Low Packed Single-Precision Floating-Point Values
func VUNPCKLPSm128float32(X1 []float32, X2 []float32, X3 []float32)

// go:noescape

// VXORPDm128byte Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func VXORPDm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VXORPDm128float64 Bitwise Logical XOR of Packed Double Precision Floating-Point Values
func VXORPDm128float64(X1 []float64, X2 []float64, X3 []float64)

// go:noescape

// VXORPSm128byte Bitwise Logical XOR of Packed Single Precision Floating-Point Values
func VXORPSm128byte(X1 []byte, X2 []byte, X3 []byte)

// go:noescape

// VXORPSm128float32 Bitwise Logical XOR of Packed Single Precision Floating-Point Values
func VXORPSm128float32(X1 []float32, X2 []float32, X3 []float32)
