package sse2

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

func TestADDPDm128byte(t *testing.T) {
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
	ADDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDPD, a=%v, b=%v", a, b)
	}
}

func TestADDPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ADDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDPD, a=%v, b=%v", a, b)
	}
}

func TestADDSDm64byte(t *testing.T) {
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
	ADDSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSD, a=%v, b=%v", a, b)
	}
}

func TestADDSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ADDSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSD, a=%v, b=%v", a, b)
	}
}

func TestANDNPDm128byte(t *testing.T) {
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
	ANDNPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ANDNPD, a=%v, b=%v", a, b)
	}
}

func TestANDNPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ANDNPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ANDNPD, a=%v, b=%v", a, b)
	}
}

func TestANDPDm128byte(t *testing.T) {
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
	ANDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ANDPD, a=%v, b=%v", a, b)
	}
}

func TestANDPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ANDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ANDPD, a=%v, b=%v", a, b)
	}
}

func TestCOMISDm64byte(t *testing.T) {
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
	COMISDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("COMISD, a=%v, b=%v", a, b)
	}
}

func TestCOMISDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	COMISDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("COMISD, a=%v, b=%v", a, b)
	}
}

func TestDIVPDm128byte(t *testing.T) {
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
	DIVPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("DIVPD, a=%v, b=%v", a, b)
	}
}

func TestDIVPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	DIVPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("DIVPD, a=%v, b=%v", a, b)
	}
}

func TestDIVSDm64byte(t *testing.T) {
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
	DIVSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("DIVSD, a=%v, b=%v", a, b)
	}
}

func TestDIVSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	DIVSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("DIVSD, a=%v, b=%v", a, b)
	}
}

func TestMASKMOVDQUbyte(t *testing.T) {
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
	MASKMOVDQUbyte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MASKMOVDQU, a=%v, b=%v", a, b)
	}
}

func TestMAXPDm128byte(t *testing.T) {
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
	MAXPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MAXPD, a=%v, b=%v", a, b)
	}
}

func TestMAXPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MAXPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MAXPD, a=%v, b=%v", a, b)
	}
}

func TestMAXSDm64byte(t *testing.T) {
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
	MAXSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MAXSD, a=%v, b=%v", a, b)
	}
}

func TestMAXSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MAXSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MAXSD, a=%v, b=%v", a, b)
	}
}

func TestMINPDm128byte(t *testing.T) {
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
	MINPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MINPD, a=%v, b=%v", a, b)
	}
}

func TestMINPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MINPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MINPD, a=%v, b=%v", a, b)
	}
}

func TestMINSDm64byte(t *testing.T) {
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
	MINSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MINSD, a=%v, b=%v", a, b)
	}
}

func TestMINSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MINSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MINSD, a=%v, b=%v", a, b)
	}
}

func TestMULPDm128byte(t *testing.T) {
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
	MULPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MULPD, a=%v, b=%v", a, b)
	}
}

func TestMULPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MULPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MULPD, a=%v, b=%v", a, b)
	}
}

func TestMULSDm64byte(t *testing.T) {
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
	MULSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MULSD, a=%v, b=%v", a, b)
	}
}

func TestMULSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	MULSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("MULSD, a=%v, b=%v", a, b)
	}
}

func TestORPDm128byte(t *testing.T) {
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
	ORPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ORPD, a=%v, b=%v", a, b)
	}
}

func TestORPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ORPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ORPD, a=%v, b=%v", a, b)
	}
}

func TestPACKSSDWm128byte(t *testing.T) {
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
	PACKSSDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PACKSSDW, a=%v, b=%v", a, b)
	}
}

func TestPACKSSWBm128byte(t *testing.T) {
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
	PACKSSWBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PACKSSWB, a=%v, b=%v", a, b)
	}
}

func TestPACKUSWBm128byte(t *testing.T) {
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
	PACKUSWBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PACKUSWB, a=%v, b=%v", a, b)
	}
}

func TestPADDBm128byte(t *testing.T) {
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
	PADDBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDB, a=%v, b=%v", a, b)
	}
}

func TestPADDBm128int8(t *testing.T) {
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
	PADDBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDB, a=%v, b=%v", a, b)
	}
}

func TestPADDDm128byte(t *testing.T) {
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
	PADDDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDD, a=%v, b=%v", a, b)
	}
}

func TestPADDDm128int32(t *testing.T) {
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
	PADDDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDD, a=%v, b=%v", a, b)
	}
}

func TestPADDQm128byte(t *testing.T) {
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
	PADDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDQ, a=%v, b=%v", a, b)
	}
}

func TestPADDQm128int64(t *testing.T) {
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
	PADDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDQ, a=%v, b=%v", a, b)
	}
}

func TestPADDSBm128byte(t *testing.T) {
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
	PADDSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDSB, a=%v, b=%v", a, b)
	}
}

func TestPADDSBm128int8(t *testing.T) {
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
	PADDSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDSB, a=%v, b=%v", a, b)
	}
}

func TestPADDSWm128byte(t *testing.T) {
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
	PADDSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDSW, a=%v, b=%v", a, b)
	}
}

func TestPADDSWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PADDSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDSW, a=%v, b=%v", a, b)
	}
}

func TestPADDUSBm128byte(t *testing.T) {
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
	PADDUSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDUSB, a=%v, b=%v", a, b)
	}
}

func TestPADDUSBm128uint8(t *testing.T) {
	a := make([]uint8, 64)
	aT := make([]uint8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint8, 64)
	bT := make([]uint8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PADDUSBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDUSB, a=%v, b=%v", a, b)
	}
}

func TestPADDUSWm128byte(t *testing.T) {
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
	PADDUSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDUSW, a=%v, b=%v", a, b)
	}
}

func TestPADDUSWm128uint16(t *testing.T) {
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
	PADDUSWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDUSW, a=%v, b=%v", a, b)
	}
}

func TestPADDWm128byte(t *testing.T) {
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
	PADDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDW, a=%v, b=%v", a, b)
	}
}

func TestPADDWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PADDWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PADDW, a=%v, b=%v", a, b)
	}
}

func TestPANDm128byte(t *testing.T) {
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
	PANDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PAND, a=%v, b=%v", a, b)
	}
}

func TestPANDNm128byte(t *testing.T) {
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
	PANDNm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PANDN, a=%v, b=%v", a, b)
	}
}

func TestPAVGBm128byte(t *testing.T) {
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
	PAVGBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PAVGB, a=%v, b=%v", a, b)
	}
}

func TestPAVGBm128int8(t *testing.T) {
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
	PAVGBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PAVGB, a=%v, b=%v", a, b)
	}
}

func TestPAVGWm128byte(t *testing.T) {
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
	PAVGWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PAVGW, a=%v, b=%v", a, b)
	}
}

func TestPAVGWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PAVGWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PAVGW, a=%v, b=%v", a, b)
	}
}

func TestPCMPEQBm128byte(t *testing.T) {
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
	PCMPEQBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPEQB, a=%v, b=%v", a, b)
	}
}

func TestPCMPEQDm128byte(t *testing.T) {
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
	PCMPEQDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPEQD, a=%v, b=%v", a, b)
	}
}

func TestPCMPEQWm128byte(t *testing.T) {
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
	PCMPEQWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPEQW, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTBm128byte(t *testing.T) {
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
	PCMPGTBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTB, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTBm128int8(t *testing.T) {
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
	PCMPGTBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTB, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTDm128byte(t *testing.T) {
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
	PCMPGTDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTD, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTDm128int32(t *testing.T) {
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
	PCMPGTDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTD, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTWm128byte(t *testing.T) {
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
	PCMPGTWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTW, a=%v, b=%v", a, b)
	}
}

func TestPCMPGTWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PCMPGTWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PCMPGTW, a=%v, b=%v", a, b)
	}
}

func TestPMADDWDm128byte(t *testing.T) {
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
	PMADDWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMADDWD, a=%v, b=%v", a, b)
	}
}

func TestPMADDWDm128int32(t *testing.T) {
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
	PMADDWDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMADDWD, a=%v, b=%v", a, b)
	}
}

func TestPMAXSWm128byte(t *testing.T) {
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
	PMAXSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSW, a=%v, b=%v", a, b)
	}
}

func TestPMAXSWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXSW, a=%v, b=%v", a, b)
	}
}

func TestPMAXUBm128byte(t *testing.T) {
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
	PMAXUBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUB, a=%v, b=%v", a, b)
	}
}

func TestPMAXUBm128uint8(t *testing.T) {
	a := make([]uint8, 64)
	aT := make([]uint8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint8, 64)
	bT := make([]uint8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMAXUBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMAXUB, a=%v, b=%v", a, b)
	}
}

func TestPMINSWm128byte(t *testing.T) {
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
	PMINSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSW, a=%v, b=%v", a, b)
	}
}

func TestPMINSWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINSW, a=%v, b=%v", a, b)
	}
}

func TestPMINUBm128byte(t *testing.T) {
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
	PMINUBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUB, a=%v, b=%v", a, b)
	}
}

func TestPMINUBm128uint8(t *testing.T) {
	a := make([]uint8, 64)
	aT := make([]uint8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint8, 64)
	bT := make([]uint8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMINUBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMINUB, a=%v, b=%v", a, b)
	}
}

func TestPMULHUWm128byte(t *testing.T) {
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
	PMULHUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULHUW, a=%v, b=%v", a, b)
	}
}

func TestPMULHUWm128uint16(t *testing.T) {
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
	PMULHUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULHUW, a=%v, b=%v", a, b)
	}
}

func TestPMULHWm128byte(t *testing.T) {
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
	PMULHWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULHW, a=%v, b=%v", a, b)
	}
}

func TestPMULHWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULHWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULHW, a=%v, b=%v", a, b)
	}
}

func TestPMULLWm128byte(t *testing.T) {
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
	PMULLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULLW, a=%v, b=%v", a, b)
	}
}

func TestPMULLWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PMULLWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULLW, a=%v, b=%v", a, b)
	}
}

func TestPMULUDQm128byte(t *testing.T) {
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
	PMULUDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULUDQ, a=%v, b=%v", a, b)
	}
}

func TestPMULUDQm128int64(t *testing.T) {
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
	PMULUDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULUDQ, a=%v, b=%v", a, b)
	}
}

func TestPORm128byte(t *testing.T) {
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
	PORm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("POR, a=%v, b=%v", a, b)
	}
}

func TestPSADBWm128byte(t *testing.T) {
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
	PSADBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSADBW, a=%v, b=%v", a, b)
	}
}

func TestPSLLDm128byte(t *testing.T) {
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
	PSLLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSLLD, a=%v, b=%v", a, b)
	}
}

func TestPSLLQm128byte(t *testing.T) {
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
	PSLLQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSLLQ, a=%v, b=%v", a, b)
	}
}

func TestPSLLWm128byte(t *testing.T) {
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
	PSLLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSLLW, a=%v, b=%v", a, b)
	}
}

func TestPSRADm128byte(t *testing.T) {
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
	PSRADm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSRAD, a=%v, b=%v", a, b)
	}
}

func TestPSRAWm128byte(t *testing.T) {
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
	PSRAWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSRAW, a=%v, b=%v", a, b)
	}
}

func TestPSRLDm128byte(t *testing.T) {
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
	PSRLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSRLD, a=%v, b=%v", a, b)
	}
}

func TestPSRLQm128byte(t *testing.T) {
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
	PSRLQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSRLQ, a=%v, b=%v", a, b)
	}
}

func TestPSRLWm128byte(t *testing.T) {
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
	PSRLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSRLW, a=%v, b=%v", a, b)
	}
}

func TestPSUBBm128byte(t *testing.T) {
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
	PSUBBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBB, a=%v, b=%v", a, b)
	}
}

func TestPSUBBm128int8(t *testing.T) {
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
	PSUBBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBB, a=%v, b=%v", a, b)
	}
}

func TestPSUBDm128byte(t *testing.T) {
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
	PSUBDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBD, a=%v, b=%v", a, b)
	}
}

func TestPSUBDm128int32(t *testing.T) {
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
	PSUBDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBD, a=%v, b=%v", a, b)
	}
}

func TestPSUBQm128byte(t *testing.T) {
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
	PSUBQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBQ, a=%v, b=%v", a, b)
	}
}

func TestPSUBQm128int64(t *testing.T) {
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
	PSUBQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBQ, a=%v, b=%v", a, b)
	}
}

func TestPSUBSBm128byte(t *testing.T) {
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
	PSUBSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBSB, a=%v, b=%v", a, b)
	}
}

func TestPSUBSBm128int8(t *testing.T) {
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
	PSUBSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBSB, a=%v, b=%v", a, b)
	}
}

func TestPSUBSWm128byte(t *testing.T) {
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
	PSUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBSW, a=%v, b=%v", a, b)
	}
}

func TestPSUBSWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PSUBSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBSW, a=%v, b=%v", a, b)
	}
}

func TestPSUBUSBm128byte(t *testing.T) {
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
	PSUBUSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBUSB, a=%v, b=%v", a, b)
	}
}

func TestPSUBUSBm128uint8(t *testing.T) {
	a := make([]uint8, 64)
	aT := make([]uint8, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]uint8, 64)
	bT := make([]uint8, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PSUBUSBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBUSB, a=%v, b=%v", a, b)
	}
}

func TestPSUBUSWm128byte(t *testing.T) {
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
	PSUBUSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBUSW, a=%v, b=%v", a, b)
	}
}

func TestPSUBUSWm128uint16(t *testing.T) {
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
	PSUBUSWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBUSW, a=%v, b=%v", a, b)
	}
}

func TestPSUBWm128byte(t *testing.T) {
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
	PSUBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBW, a=%v, b=%v", a, b)
	}
}

func TestPSUBWm128int16(t *testing.T) {
	a := make([]int16, 64)
	aT := make([]int16, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]int16, 64)
	bT := make([]int16, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	PSUBWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSUBW, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKHBWm128byte(t *testing.T) {
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
	PUNPCKHBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKHBW, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKHDQm128byte(t *testing.T) {
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
	PUNPCKHDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKHDQ, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKHQDQm128byte(t *testing.T) {
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
	PUNPCKHQDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKHQDQ, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKHWDm128byte(t *testing.T) {
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
	PUNPCKHWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKHWD, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKLBWm128byte(t *testing.T) {
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
	PUNPCKLBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKLBW, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKLDQm128byte(t *testing.T) {
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
	PUNPCKLDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKLDQ, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKLQDQm128byte(t *testing.T) {
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
	PUNPCKLQDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKLQDQ, a=%v, b=%v", a, b)
	}
}

func TestPUNPCKLWDm128byte(t *testing.T) {
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
	PUNPCKLWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PUNPCKLWD, a=%v, b=%v", a, b)
	}
}

func TestPXORm128byte(t *testing.T) {
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
	PXORm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PXOR, a=%v, b=%v", a, b)
	}
}

func TestSQRTPDm128byte(t *testing.T) {
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
	SQRTPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SQRTPD, a=%v, b=%v", a, b)
	}
}

func TestSQRTPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	SQRTPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SQRTPD, a=%v, b=%v", a, b)
	}
}

func TestSQRTSDm64byte(t *testing.T) {
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
	SQRTSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SQRTSD, a=%v, b=%v", a, b)
	}
}

func TestSQRTSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	SQRTSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SQRTSD, a=%v, b=%v", a, b)
	}
}

func TestSUBPDm128byte(t *testing.T) {
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
	SUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SUBPD, a=%v, b=%v", a, b)
	}
}

func TestSUBPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	SUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SUBPD, a=%v, b=%v", a, b)
	}
}

func TestSUBSDm64byte(t *testing.T) {
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
	SUBSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SUBSD, a=%v, b=%v", a, b)
	}
}

func TestSUBSDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	SUBSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("SUBSD, a=%v, b=%v", a, b)
	}
}

func TestUCOMISDm64byte(t *testing.T) {
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
	UCOMISDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UCOMISD, a=%v, b=%v", a, b)
	}
}

func TestUCOMISDm64float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	UCOMISDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UCOMISD, a=%v, b=%v", a, b)
	}
}

func TestUNPCKHPDm128byte(t *testing.T) {
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
	UNPCKHPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UNPCKHPD, a=%v, b=%v", a, b)
	}
}

func TestUNPCKHPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	UNPCKHPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UNPCKHPD, a=%v, b=%v", a, b)
	}
}

func TestUNPCKLPDm128byte(t *testing.T) {
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
	UNPCKLPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UNPCKLPD, a=%v, b=%v", a, b)
	}
}

func TestUNPCKLPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	UNPCKLPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("UNPCKLPD, a=%v, b=%v", a, b)
	}
}

func TestXORPDm128byte(t *testing.T) {
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
	XORPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("XORPD, a=%v, b=%v", a, b)
	}
}

func TestXORPDm128float64(t *testing.T) {
	a := make([]float64, 64)
	aT := make([]float64, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float64, 64)
	bT := make([]float64, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	XORPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("XORPD, a=%v, b=%v", a, b)
	}
}
