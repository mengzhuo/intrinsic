package ssse3

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

func TestPABSBm128byte(t *testing.T) {
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
	PABSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSB, a=%v, b=%v", a, b)
	}
}

func TestPABSBm128int8(t *testing.T) {
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
	PABSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSB, a=%v, b=%v", a, b)
	}
}

func TestPABSDm128byte(t *testing.T) {
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
	PABSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSD, a=%v, b=%v", a, b)
	}
}

func TestPABSDm128int32(t *testing.T) {
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
	PABSDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSD, a=%v, b=%v", a, b)
	}
}

func TestPABSWm128byte(t *testing.T) {
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
	PABSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSW, a=%v, b=%v", a, b)
	}
}

func TestPABSWm128int16(t *testing.T) {
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
	PABSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PABSW, a=%v, b=%v", a, b)
	}
}

func TestPHADDDm128byte(t *testing.T) {
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
	PHADDDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHADDD, a=%v, b=%v", a, b)
	}
}

func TestPHADDSWm128byte(t *testing.T) {
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
	PHADDSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHADDSW, a=%v, b=%v", a, b)
	}
}

func TestPHADDWm128byte(t *testing.T) {
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
	PHADDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHADDW, a=%v, b=%v", a, b)
	}
}

func TestPHSUBDm128byte(t *testing.T) {
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
	PHSUBDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHSUBD, a=%v, b=%v", a, b)
	}
}

func TestPHSUBSWm128byte(t *testing.T) {
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
	PHSUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHSUBSW, a=%v, b=%v", a, b)
	}
}

func TestPHSUBWm128byte(t *testing.T) {
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
	PHSUBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PHSUBW, a=%v, b=%v", a, b)
	}
}

func TestPMADDUBSWm128byte(t *testing.T) {
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
	PMADDUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMADDUBSW, a=%v, b=%v", a, b)
	}
}

func TestPMULHRSWm128byte(t *testing.T) {
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
	PMULHRSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PMULHRSW, a=%v, b=%v", a, b)
	}
}

func TestPSHUFBm128byte(t *testing.T) {
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
	PSHUFBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSHUFB, a=%v, b=%v", a, b)
	}
}

func TestPSIGNBm128byte(t *testing.T) {
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
	PSIGNBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSIGNB, a=%v, b=%v", a, b)
	}
}

func TestPSIGNDm128byte(t *testing.T) {
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
	PSIGNDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSIGND, a=%v, b=%v", a, b)
	}
}

func TestPSIGNWm128byte(t *testing.T) {
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
	PSIGNWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("PSIGNW, a=%v, b=%v", a, b)
	}
}
