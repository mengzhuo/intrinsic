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



TEXT ·ADDSUBPDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	BYTE $0x66; BYTE $0x0F; BYTE $0xD0;  BYTE $0xca // $0xca = X2, X1
	RETX1X2
	

TEXT ·ADDSUBPDm128float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	BYTE $0x66; BYTE $0x0F; BYTE $0xD0;  BYTE $0xca // $0xca = X2, X1
	RETX1X2
	


TEXT ·ADDSUBPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	BYTE $0xF2; BYTE $0x0F; BYTE $0xD0;  BYTE $0xca // $0xca = X2, X1
	RETX1X2
	

TEXT ·ADDSUBPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	BYTE $0xF2; BYTE $0x0F; BYTE $0xD0;  BYTE $0xca // $0xca = X2, X1
	RETX1X2
	


TEXT ·HADDPDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	HADDPD X2, X1
	RETX1X2
	

TEXT ·HADDPDm128float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	HADDPD X2, X1
	RETX1X2
	


TEXT ·HADDPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	HADDPS X2, X1
	RETX1X2
	

TEXT ·HADDPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	HADDPS X2, X1
	RETX1X2
	


TEXT ·HSUBPDm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	HSUBPD X2, X1
	RETX1X2
	

TEXT ·HSUBPDm128float64(SB),NOSPLIT,$0-48
	FPTOX1X2
	HSUBPD X2, X1
	RETX1X2
	


TEXT ·HSUBPSm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	HSUBPS X2, X1
	RETX1X2
	

TEXT ·HSUBPSm128float32(SB),NOSPLIT,$0-48
	FPTOX1X2
	HSUBPS X2, X1
	RETX1X2
	


