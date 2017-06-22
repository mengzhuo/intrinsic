package sse2

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
		t.Error("Nothing changed on ADDPD")
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
		t.Error("Nothing changed on ADDPD")
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
		t.Error("Nothing changed on ADDSD")
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
		t.Error("Nothing changed on ADDSD")
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
		t.Error("Nothing changed on ANDNPD")
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
		t.Error("Nothing changed on ANDNPD")
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
		t.Error("Nothing changed on ANDPD")
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
		t.Error("Nothing changed on ANDPD")
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
		t.Error("Nothing changed on COMISD")
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
		t.Error("Nothing changed on COMISD")
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
		t.Error("Nothing changed on DIVPD")
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
		t.Error("Nothing changed on DIVPD")
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
		t.Error("Nothing changed on DIVSD")
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
		t.Error("Nothing changed on DIVSD")
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
		t.Error("Nothing changed on MASKMOVDQU")
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
		t.Error("Nothing changed on MAXPD")
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
		t.Error("Nothing changed on MAXPD")
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
		t.Error("Nothing changed on MAXSD")
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
		t.Error("Nothing changed on MAXSD")
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
		t.Error("Nothing changed on MINPD")
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
		t.Error("Nothing changed on MINPD")
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
		t.Error("Nothing changed on MINSD")
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
		t.Error("Nothing changed on MINSD")
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
		t.Error("Nothing changed on MULPD")
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
		t.Error("Nothing changed on MULPD")
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
		t.Error("Nothing changed on MULSD")
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
		t.Error("Nothing changed on MULSD")
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
		t.Error("Nothing changed on ORPD")
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
		t.Error("Nothing changed on ORPD")
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
		t.Error("Nothing changed on PACKSSDW")
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
		t.Error("Nothing changed on PACKSSWB")
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
		t.Error("Nothing changed on PACKUSWB")
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
		t.Error("Nothing changed on PADDB")
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
		t.Error("Nothing changed on PADDB")
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
		t.Error("Nothing changed on PADDD")
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
		t.Error("Nothing changed on PADDD")
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
		t.Error("Nothing changed on PADDQ")
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
		t.Error("Nothing changed on PADDQ")
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
		t.Error("Nothing changed on PADDSB")
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
		t.Error("Nothing changed on PADDSB")
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
		t.Error("Nothing changed on PADDSW")
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
		t.Error("Nothing changed on PADDSW")
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
		t.Error("Nothing changed on PADDUSB")
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
		t.Error("Nothing changed on PADDUSB")
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
		t.Error("Nothing changed on PADDUSW")
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
		t.Error("Nothing changed on PADDUSW")
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
		t.Error("Nothing changed on PADDW")
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
		t.Error("Nothing changed on PADDW")
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
		t.Error("Nothing changed on PAND")
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
		t.Error("Nothing changed on PANDN")
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
		t.Error("Nothing changed on PAVGB")
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
		t.Error("Nothing changed on PAVGB")
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
		t.Error("Nothing changed on PAVGW")
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
		t.Error("Nothing changed on PAVGW")
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
		t.Error("Nothing changed on PCMPEQB")
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
		t.Error("Nothing changed on PCMPEQD")
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
		t.Error("Nothing changed on PCMPEQW")
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
		t.Error("Nothing changed on PCMPGTB")
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
		t.Error("Nothing changed on PCMPGTB")
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
		t.Error("Nothing changed on PCMPGTD")
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
		t.Error("Nothing changed on PCMPGTD")
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
		t.Error("Nothing changed on PCMPGTW")
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
		t.Error("Nothing changed on PCMPGTW")
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
		t.Error("Nothing changed on PMADDWD")
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
		t.Error("Nothing changed on PMADDWD")
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
		t.Error("Nothing changed on PMAXSW")
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
		t.Error("Nothing changed on PMAXSW")
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
		t.Error("Nothing changed on PMAXUB")
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
		t.Error("Nothing changed on PMAXUB")
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
		t.Error("Nothing changed on PMINSW")
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
		t.Error("Nothing changed on PMINSW")
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
		t.Error("Nothing changed on PMINUB")
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
		t.Error("Nothing changed on PMINUB")
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
		t.Error("Nothing changed on PMULHUW")
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
		t.Error("Nothing changed on PMULHUW")
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
		t.Error("Nothing changed on PMULHW")
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
		t.Error("Nothing changed on PMULHW")
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
		t.Error("Nothing changed on PMULLW")
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
		t.Error("Nothing changed on PMULLW")
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
		t.Error("Nothing changed on PMULUDQ")
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
		t.Error("Nothing changed on PMULUDQ")
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
		t.Error("Nothing changed on POR")
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
		t.Error("Nothing changed on PSADBW")
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
		t.Error("Nothing changed on PSLLD")
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
		t.Error("Nothing changed on PSLLQ")
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
		t.Error("Nothing changed on PSLLW")
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
		t.Error("Nothing changed on PSRAD")
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
		t.Error("Nothing changed on PSRAW")
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
		t.Error("Nothing changed on PSRLD")
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
		t.Error("Nothing changed on PSRLQ")
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
		t.Error("Nothing changed on PSRLW")
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
		t.Error("Nothing changed on PSUBB")
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
		t.Error("Nothing changed on PSUBB")
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
		t.Error("Nothing changed on PSUBD")
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
		t.Error("Nothing changed on PSUBD")
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
		t.Error("Nothing changed on PSUBQ")
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
		t.Error("Nothing changed on PSUBQ")
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
		t.Error("Nothing changed on PSUBSB")
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
		t.Error("Nothing changed on PSUBSB")
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
		t.Error("Nothing changed on PSUBSW")
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
		t.Error("Nothing changed on PSUBSW")
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
		t.Error("Nothing changed on PSUBUSB")
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
		t.Error("Nothing changed on PSUBUSB")
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
		t.Error("Nothing changed on PSUBUSW")
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
		t.Error("Nothing changed on PSUBUSW")
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
		t.Error("Nothing changed on PSUBW")
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
		t.Error("Nothing changed on PSUBW")
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
		t.Error("Nothing changed on PUNPCKHBW")
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
		t.Error("Nothing changed on PUNPCKHDQ")
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
		t.Error("Nothing changed on PUNPCKHQDQ")
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
		t.Error("Nothing changed on PUNPCKHWD")
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
		t.Error("Nothing changed on PUNPCKLBW")
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
		t.Error("Nothing changed on PUNPCKLDQ")
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
		t.Error("Nothing changed on PUNPCKLQDQ")
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
		t.Error("Nothing changed on PUNPCKLWD")
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
		t.Error("Nothing changed on PXOR")
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
		t.Error("Nothing changed on SQRTPD")
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
		t.Error("Nothing changed on SQRTPD")
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
		t.Error("Nothing changed on SQRTSD")
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
		t.Error("Nothing changed on SQRTSD")
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
		t.Error("Nothing changed on SUBPD")
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
		t.Error("Nothing changed on SUBPD")
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
		t.Error("Nothing changed on SUBSD")
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
		t.Error("Nothing changed on SUBSD")
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
		t.Error("Nothing changed on UCOMISD")
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
		t.Error("Nothing changed on UCOMISD")
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
		t.Error("Nothing changed on UNPCKHPD")
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
		t.Error("Nothing changed on UNPCKHPD")
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
		t.Error("Nothing changed on UNPCKLPD")
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
		t.Error("Nothing changed on UNPCKLPD")
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
		t.Error("Nothing changed on XORPD")
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
		t.Error("Nothing changed on XORPD")
	}
}
