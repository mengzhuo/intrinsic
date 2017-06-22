#include "textflag.h"

#define FPTOX1X2 \
	MOVQ a+0(FP), SI;\
	MOVQ b+24(FP), DI;\
	MOVOU (SI), X1;\
	MOVOU (DI), X2;\

#define RETX1X2 \
	MOVOU X1, (SI);\
	MOVOU X2, (DI);\
	RET;\

#define FPTOY1Y2 \
	MOVQ a+0(FP), SI;\
	MOVQ b+24(FP), DI;\
	MOVOU (SI), Y1;\
	MOVOU (DI), Y2;\

#define FPTOX1X2X3 \
	MOVQ a+0(FP), SI;\
	MOVQ b+24(FP), DI;\
	MOVOU (SI), X1;\
	MOVOU (DI), X2;\
	MOVQ c+48(FP), DI;\
	MOVOU (DI), X3;\

#define RETY1Y2 \
	MOVOU Y1, (SI);\
	MOVOU Y2, (DI);\
	RET;\



TEXT ·VADDPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VADDPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VADDSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VADDSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VADDSUBPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDSUBPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VADDSUBPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VADDSUBPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VADDSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VANDNPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDNPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VANDNPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDNPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VANDNPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDNPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VANDNPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDNPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VANDPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VANDPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VANDPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VANDPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VANDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VCOMISDm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VCOMISD X2, X1
	RETX1X2
	

TEXT ·VCOMISDm64float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	VCOMISD X2, X1
	RETX1X2
	


TEXT ·VCOMISSm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VCOMISS X2, X1
	RETX1X2
	

TEXT ·VCOMISSm32float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	VCOMISS X2, X1
	RETX1X2
	


TEXT ·VDIVPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VDIVPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VDIVPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VDIVPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VDIVSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VDIVSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VDIVSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VDIVSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VDIVSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VHADDPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHADDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VHADDPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHADDPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VHADDPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHADDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VHADDPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHADDPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VHSUBPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VHSUBPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VHSUBPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VHSUBPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VHSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMASKMOVDQUbyte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VMASKMOVDQU X2, X1
	RETX1X2
	


TEXT ·VMAXPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMAXPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMAXPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMAXPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMAXSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMAXSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMAXSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMAXSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMAXSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMINPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMINPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMINPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMINPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMINSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMINSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMINSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMINSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMINSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMULPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMULPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMULPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMULPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMULSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMULSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VMULSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VMULSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VMULSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VORPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VORPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VORPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VORPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VORPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VORPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VORPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VORPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPABSBm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPABSB X2, X1
	RETX1X2
	


TEXT ·VPABSDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPABSD X2, X1
	RETX1X2
	


TEXT ·VPABSWm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPABSW X2, X1
	RETX1X2
	


TEXT ·VPACKSSDWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPACKSSDW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPACKSSWBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPACKSSWB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPACKUSDWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPACKUSDW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPACKUSWBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPACKUSWB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDQm128int64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDSBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDSWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDUSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDUSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDUSBm128uint8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDUSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDUSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDUSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDUSWm128uint16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDUSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPADDWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPADDWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPADDW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPANDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPAND X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPANDNm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPANDN X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPAVGBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPAVGB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPAVGBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPAVGB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPAVGWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPAVGW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPAVGWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPAVGW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPEQBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPEQB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPEQDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPEQD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPEQQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPEQQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPEQWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPEQW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPGTBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPCMPGTBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPGTDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPCMPGTDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPGTQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPCMPGTWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPCMPGTWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPCMPGTW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPERMILPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPERMILPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPERMILPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPERMILPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPERMILPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPERMILPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPERMILPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPERMILPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHADDDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHADDD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHADDSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHADDSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHADDWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHADDW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHMINPOSUWm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPHMINPOSUW X2, X1
	RETX1X2
	


TEXT ·VPHSUBDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHSUBD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHSUBSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHSUBSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPHSUBWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPHSUBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMADDUBSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMADDUBSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMADDWDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMADDWD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMADDWDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMADDWD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXSBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXSDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXSDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXSWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXUBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXUBm128uint8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXUDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXUDm128uint32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMAXUWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMAXUWm128uint16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMAXUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINSBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINSDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINSDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINSWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINUBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINUBm128uint8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINUDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINUDm128uint32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMINUWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMINUWm128uint16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMINUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMOVSXBDm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXBD X2, X1
	RETX1X2
	


TEXT ·VPMOVSXBQm16byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXBQ X2, X1
	RETX1X2
	


TEXT ·VPMOVSXBWm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXBW X2, X1
	RETX1X2
	


TEXT ·VPMOVSXDQm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXDQ X2, X1
	RETX1X2
	


TEXT ·VPMOVSXWDm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXWD X2, X1
	RETX1X2
	


TEXT ·VPMOVSXWQm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVSXWQ X2, X1
	RETX1X2
	


TEXT ·VPMOVZXBDm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXBD X2, X1
	RETX1X2
	


TEXT ·VPMOVZXBQm16byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXBQ X2, X1
	RETX1X2
	


TEXT ·VPMOVZXBWm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXBW X2, X1
	RETX1X2
	


TEXT ·VPMOVZXDQm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXDQ X2, X1
	RETX1X2
	


TEXT ·VPMOVZXWDm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXWD X2, X1
	RETX1X2
	


TEXT ·VPMOVZXWQm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPMOVZXWQ X2, X1
	RETX1X2
	


TEXT ·VPMULDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULDQm128int64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULHRSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULHRSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULHUWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULHUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULHUWm128uint16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULHUW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULHWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULHW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULHWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULHW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULLDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULLD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULLDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULLD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULLWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULLW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULLWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULLW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPMULUDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULUDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPMULUDQm128int64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPMULUDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPORm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPOR X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSADBWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSADBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSHUFBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSHUFB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSIGNBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSIGNB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSIGNDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSIGND X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSIGNWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSIGNW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSLLDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSLLD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSLLQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSLLQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSLLWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSLLW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSRADm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSRAD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSRAWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSRAW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSRLDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSRLD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSRLQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSRLQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSRLWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSRLW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBDm128int32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBQm128int64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBSBm128int8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBSWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBUSBm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBUSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBUSBm128uint8(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBUSB X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBUSWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBUSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBUSWm128uint16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBUSW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPSUBWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VPSUBWm128int16(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPSUBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPTESTm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VPTEST X2, X1
	RETX1X2
	


TEXT ·VPTESTm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VPTEST Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VPUNPCKHBWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKHBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKHDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKHDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKHQDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKHQDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKHWDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKHWD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKLBWm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKLBW X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKLDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKLDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKLQDQm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKLQDQ X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPUNPCKLWDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPUNPCKLWD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VPXORm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VPXOR X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VRCPPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VRCPPS X2, X1
	RETX1X2
	

TEXT ·VRCPPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	VRCPPS X2, X1
	RETX1X2
	


TEXT ·VRCPPSm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VRCPPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	

TEXT ·VRCPPSm256float32(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VRCPPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VRCPSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VRCPSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VRCPSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VRCPSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VRSQRTPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VRSQRTPS X2, X1
	RETX1X2
	

TEXT ·VRSQRTPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	VRSQRTPS X2, X1
	RETX1X2
	


TEXT ·VRSQRTPSm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VRSQRTPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	

TEXT ·VRSQRTPSm256float32(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VRSQRTPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VRSQRTSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VRSQRTSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VRSQRTSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VRSQRTSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSQRTPDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VSQRTPD X2, X1
	RETX1X2
	

TEXT ·VSQRTPDm128float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	VSQRTPD X2, X1
	RETX1X2
	


TEXT ·VSQRTPDm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VSQRTPD Y2, Y1
	MOVOU Y1, (SI)
	RET
	

TEXT ·VSQRTPDm256float64(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VSQRTPD Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VSQRTPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VSQRTPS X2, X1
	RETX1X2
	

TEXT ·VSQRTPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	VSQRTPS X2, X1
	RETX1X2
	


TEXT ·VSQRTPSm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VSQRTPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	

TEXT ·VSQRTPSm256float32(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VSQRTPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VSQRTSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSQRTSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VSQRTSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSQRTSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSQRTSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSQRTSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSUBPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VSUBPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSUBPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VSUBPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSUBSDm64byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VSUBSDm64float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBSD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VSUBSSm32byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VSUBSSm32float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VSUBSS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VTESTPDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VTESTPD X2, X1
	RETX1X2
	


TEXT ·VTESTPDm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VTESTPD Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VTESTPSm256byte(SB),NOSPLIT,$0-48
	
	FPTOY1Y2
	VTESTPS Y2, Y1
	MOVOU Y1, (SI)
	RET
	


TEXT ·VTESTPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VTESTPS X2, X1
	RETX1X2
	


TEXT ·VUCOMISDm64byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VUCOMISD X2, X1
	RETX1X2
	

TEXT ·VUCOMISDm64float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	VUCOMISD X2, X1
	RETX1X2
	


TEXT ·VUCOMISSm32byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	VUCOMISS X2, X1
	RETX1X2
	

TEXT ·VUCOMISSm32float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	VUCOMISS X2, X1
	RETX1X2
	


TEXT ·VUNPCKHPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKHPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VUNPCKHPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKHPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VUNPCKHPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKHPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VUNPCKHPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKHPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VUNPCKLPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKLPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VUNPCKLPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKLPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VUNPCKLPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKLPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VUNPCKLPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VUNPCKLPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VXORPDm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VXORPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VXORPDm128float64(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VXORPD X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


TEXT ·VXORPSm128byte(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VXORPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	

TEXT ·VXORPSm128float32(SB),NOSPLIT,$0-72
	
	FPTOX1X2X3
	VXORPS X3, X2, X1
	MOVOU X1, (SI)
	MOVOU X2, (DI)
	RET
	


