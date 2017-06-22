package sse3

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

func TestADDSUBPDm128byte(t *testing.T) {
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
	ADDSUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSUBPD, a=%v, b=%v", a, b)
	}
}

func TestADDSUBPDm128float64(t *testing.T) {
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
	ADDSUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSUBPD, a=%v, b=%v", a, b)
	}
}

func TestADDSUBPSm128byte(t *testing.T) {
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
	ADDSUBPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSUBPS, a=%v, b=%v", a, b)
	}
}

func TestADDSUBPSm128float32(t *testing.T) {
	a := make([]float32, 64)
	aT := make([]float32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float32, 64)
	bT := make([]float32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	ADDSUBPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("ADDSUBPS, a=%v, b=%v", a, b)
	}
}

func TestHADDPDm128byte(t *testing.T) {
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
	HADDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HADDPD, a=%v, b=%v", a, b)
	}
}

func TestHADDPDm128float64(t *testing.T) {
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
	HADDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HADDPD, a=%v, b=%v", a, b)
	}
}

func TestHADDPSm128byte(t *testing.T) {
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
	HADDPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HADDPS, a=%v, b=%v", a, b)
	}
}

func TestHADDPSm128float32(t *testing.T) {
	a := make([]float32, 64)
	aT := make([]float32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float32, 64)
	bT := make([]float32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	HADDPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HADDPS, a=%v, b=%v", a, b)
	}
}

func TestHSUBPDm128byte(t *testing.T) {
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
	HSUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HSUBPD, a=%v, b=%v", a, b)
	}
}

func TestHSUBPDm128float64(t *testing.T) {
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
	HSUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HSUBPD, a=%v, b=%v", a, b)
	}
}

func TestHSUBPSm128byte(t *testing.T) {
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
	HSUBPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HSUBPS, a=%v, b=%v", a, b)
	}
}

func TestHSUBPSm128float32(t *testing.T) {
	a := make([]float32, 64)
	aT := make([]float32, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]float32, 64)
	bT := make([]float32, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)
	HSUBPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Logf("HSUBPS, a=%v, b=%v", a, b)
	}
}
