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



TEXT ·PCMPGTQm128byte(SB),NOSPLIT,$0-48
	FPTOX1X2
	PCMPGTQ X2, X1
	RETX1X2
	

