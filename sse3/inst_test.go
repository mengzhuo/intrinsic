package sse3

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
		t.Error("Nothing changed on ADDSUBPD")
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
		t.Error("Nothing changed on ADDSUBPD")
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
		t.Error("Nothing changed on ADDSUBPS")
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
		t.Error("Nothing changed on ADDSUBPS")
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
		t.Error("Nothing changed on HADDPD")
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
		t.Error("Nothing changed on HADDPD")
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
		t.Error("Nothing changed on HADDPS")
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
		t.Error("Nothing changed on HADDPS")
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
		t.Error("Nothing changed on HSUBPD")
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
		t.Error("Nothing changed on HSUBPD")
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
		t.Error("Nothing changed on HSUBPS")
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
		t.Error("Nothing changed on HSUBPS")
	}
}
