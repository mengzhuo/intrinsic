package sse41

import (
	"strings"
	"testing"
)

func sh(s string) bool {

	sss := []string{"Low", "High", "Test"}

	for _, cmp := range sss {
		if strings.Index(s, cmp) != -1 {
			return true
		}
	}
	return false
}

func TestPACKUSDWm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PACKUSDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PACKUSDW, a=%v, b=%v", a, b)
	}
}

func TestPCMPEQQm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PCMPEQQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPEQQ, a=%v, b=%v", a, b)
	}
}

func TestPHMINPOSUWm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PHMINPOSUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHMINPOSUW, a=%v, b=%v", a, b)
	}
}

func TestPMAXSBm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSB, a=%v, b=%v", a, b)
	}
}

func TestPMAXSBm128int8(t *testing.T) {
	a := make([]int8, 64)
	aT := make([]int8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int8, 64)
	bT := make([]int8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSB, a=%v, b=%v", a, b)
	}
}

func TestPMAXSDm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSD, a=%v, b=%v", a, b)
	}
}

func TestPMAXSDm128int32(t *testing.T) {
	a := make([]int32, 64)
	aT := make([]int32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int32, 64)
	bT := make([]int32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXSDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSD, a=%v, b=%v", a, b)
	}
}

func TestPMAXUDm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXUDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUD, a=%v, b=%v", a, b)
	}
}

func TestPMAXUDm128uint32(t *testing.T) {
	a := make([]uint32, 64)
	aT := make([]uint32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint32, 64)
	bT := make([]uint32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXUDm128uint32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUD, a=%v, b=%v", a, b)
	}
}

func TestPMAXUWm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUW, a=%v, b=%v", a, b)
	}
}

func TestPMAXUWm128uint16(t *testing.T) {
	a := make([]uint16, 64)
	aT := make([]uint16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint16, 64)
	bT := make([]uint16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUW, a=%v, b=%v", a, b)
	}
}

func TestPMINSBm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSB, a=%v, b=%v", a, b)
	}
}

func TestPMINSBm128int8(t *testing.T) {
	a := make([]int8, 64)
	aT := make([]int8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int8, 64)
	bT := make([]int8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSB, a=%v, b=%v", a, b)
	}
}

func TestPMINSDm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSD, a=%v, b=%v", a, b)
	}
}

func TestPMINSDm128int32(t *testing.T) {
	a := make([]int32, 64)
	aT := make([]int32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int32, 64)
	bT := make([]int32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINSDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSD, a=%v, b=%v", a, b)
	}
}

func TestPMINUDm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINUDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUD, a=%v, b=%v", a, b)
	}
}

func TestPMINUDm128uint32(t *testing.T) {
	a := make([]uint32, 64)
	aT := make([]uint32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint32, 64)
	bT := make([]uint32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINUDm128uint32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUD, a=%v, b=%v", a, b)
	}
}

func TestPMINUWm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUW, a=%v, b=%v", a, b)
	}
}

func TestPMINUWm128uint16(t *testing.T) {
	a := make([]uint16, 64)
	aT := make([]uint16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint16, 64)
	bT := make([]uint16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUW, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXBDm32byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXBDm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXBD, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXBQm16byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXBQm16byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXBQ, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXBWm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXBWm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXBW, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXDQm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXDQm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXDQ, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXWDm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXWDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXWD, a=%v, b=%v", a, b)
	}
}

func TestPMOVSXWQm32byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVSXWQm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVSXWQ, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXBDm32byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXBDm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXBD, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXBQm16byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXBQm16byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXBQ, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXBWm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXBWm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXBW, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXDQm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXDQm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXDQ, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXWDm64byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXWDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXWD, a=%v, b=%v", a, b)
	}
}

func TestPMOVZXWQm32byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMOVZXWQm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMOVZXWQ, a=%v, b=%v", a, b)
	}
}

func TestPMULDQm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULDQ, a=%v, b=%v", a, b)
	}
}

func TestPMULDQm128int64(t *testing.T) {
	a := make([]int64, 64)
	aT := make([]int64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int64, 64)
	bT := make([]int64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULDQ, a=%v, b=%v", a, b)
	}
}

func TestPMULLDm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULLD, a=%v, b=%v", a, b)
	}
}

func TestPMULLDm128int32(t *testing.T) {
	a := make([]int32, 64)
	aT := make([]int32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int32, 64)
	bT := make([]int32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULLDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULLD, a=%v, b=%v", a, b)
	}
}

func TestPTESTm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PTESTm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PTEST, a=%v, b=%v", a, b)
	}
}
