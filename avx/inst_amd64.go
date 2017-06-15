package avx



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

// Convert Packed Doubleword Integers to Packed Double-Precision Floating-Point Values
func VCVTDQ2PDm64byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert Packed Doubleword Integers to Packed Double-Precision Floating-Point Values
func VCVTDQ2PDm64int32(X1 []int32,  X2 []int32)


// go:noescape

// Convert Packed Doubleword Integers to Packed Single-Precision Floating-Point Values
func VCVTDQ2PSm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Convert Packed Doubleword Integers to Packed Single-Precision Floating-Point Values
func VCVTDQ2PSm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Convert Packed Double-Precision Floating-Point Values to Packed Doubleword Integers
func VCVTPD2DQm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert Packed Double-Precision Floating-Point Values to Packed Doubleword Integers
func VCVTPD2DQm128int64(X1 []int64,  X2 []int64)


// go:noescape

// Convert Packed Double-Precision Floating-Point Values to Packed Single-Precision Floating-Point Values
func VCVTPD2PSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert Packed Double-Precision Floating-Point Values to Packed Single-Precision Floating-Point Values
func VCVTPD2PSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTPS2DQm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTPS2DQm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTPS2DQm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTPS2DQm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Double-Precision Floating-Point Values
func VCVTPS2PDm64byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert Packed Single-Precision Floating-Point Values to Packed Double-Precision Floating-Point Values
func VCVTPS2PDm64float32(X1 []float32,  X2 []float32)


// go:noescape

// Convert with Truncation Packed Double-Precision Floating-Point Values to Packed Doubleword Integers
func VCVTTPD2DQm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert with Truncation Packed Double-Precision Floating-Point Values to Packed Doubleword Integers
func VCVTTPD2DQm128int64(X1 []int64,  X2 []int64)


// go:noescape

// Convert with Truncation Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTTPS2DQm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Convert with Truncation Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTTPS2DQm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Convert with Truncation Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTTPS2DQm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Convert with Truncation Packed Single-Precision Floating-Point Values to Packed Signed Doubleword Integer Values
func VCVTTPS2DQm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Store Selected Bytes of Double Quadword
func VMASKMOVDQUbyte(X1 []byte,  X2 []byte)


// go:noescape

// Move Aligned Packed Double-Precision Floating-Point Values
func VMOVAPDm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Move Aligned Packed Double-Precision Floating-Point Values
func VMOVAPDm128float64(X1 []float64,  X2 []float64)


// go:noescape

// Move Aligned Packed Double-Precision Floating-Point Values
func VMOVAPDm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Move Aligned Packed Double-Precision Floating-Point Values
func VMOVAPDm256float64(Y1 []float64,  Y2 []float64)


// go:noescape

// Move Aligned Packed Single-Precision Floating-Point Values
func VMOVAPSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Move Aligned Packed Single-Precision Floating-Point Values
func VMOVAPSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Move Aligned Packed Single-Precision Floating-Point Values
func VMOVAPSm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Move Aligned Packed Single-Precision Floating-Point Values
func VMOVAPSm256float32(Y1 []float32,  Y2 []float32)


// go:noescape

// Replicate Double FP Values
func VMOVDDUPm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Replicate Double FP Values
func VMOVDDUPm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Move Aligned Packed Integer Values
func VMOVDQAm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Move Aligned Packed Integer Values
func VMOVDQAm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Move Unaligned Packed Integer Values
func VMOVDQUm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Move Unaligned Packed Integer Values
func VMOVDQUm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Move Quadword
func VMOVQbyte(X1 []byte,  X2 []byte)


// go:noescape

// Move Quadword
func VMOVQm64byte(X1 []byte,  X2 []byte)


// go:noescape

// Replicate Single FP Values
func VMOVSHDUPm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Replicate Single FP Values
func VMOVSHDUPm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Replicate Single FP Values
func VMOVSLDUPm128byte(X1 []byte,  X2 []byte)


// go:noescape

// Replicate Single FP Values
func VMOVSLDUPm256byte(Y1 []byte,  Y2 []byte)


// go:noescape

// Move Unaligned Packed Double-Precision Floating-Point Values
func VMOVUPDm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Move Unaligned Packed Double-Precision Floating-Point Values
func VMOVUPDm128float64(X1 []float64,  X2 []float64)


// go:noescape

// Move Unaligned Packed Double-Precision Floating-Point Values
func VMOVUPDm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Move Unaligned Packed Double-Precision Floating-Point Values
func VMOVUPDm256float64(Y1 []float64,  Y2 []float64)


// go:noescape

// Move Unaligned Packed Single-Precision Floating-Point Values
func VMOVUPSm128byte(X1 []byte,  X2 []byte)

// go:noescape

// Move Unaligned Packed Single-Precision Floating-Point Values
func VMOVUPSm128float32(X1 []float32,  X2 []float32)


// go:noescape

// Move Unaligned Packed Single-Precision Floating-Point Values
func VMOVUPSm256byte(Y1 []byte,  Y2 []byte)

// go:noescape

// Move Unaligned Packed Single-Precision Floating-Point Values
func VMOVUPSm256float32(Y1 []float32,  Y2 []float32)


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

// Packed Horizontal Word Minimum
func VPHMINPOSUWm128byte(X1 []byte,  X2 []byte)


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

// PTEST- Logical Compare
func VPTESTm128byte(X1 []byte,  X2 []byte)


// go:noescape

// PTEST- Logical Compare
func VPTESTm256byte(Y1 []byte,  Y2 []byte)


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

