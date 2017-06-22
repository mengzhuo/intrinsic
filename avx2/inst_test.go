package avx2

func TestVBROADCASTSSbyte(t *testing.T) {
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

	VBROADCASTSSbyte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VBROADCASTSS")
	}
}

func TestVPABSBm256byte(t *testing.T) {
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

	VPABSBm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSB")
	}
}

func TestVPABSDm256byte(t *testing.T) {
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

	VPABSDm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSD")
	}
}

func TestVPABSWm256byte(t *testing.T) {
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

	VPABSWm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSW")
	}
}

func TestVPBROADCASTBm8byte(t *testing.T) {
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

	VPBROADCASTBm8byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPBROADCASTB")
	}
}

func TestVPBROADCASTDm32byte(t *testing.T) {
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

	VPBROADCASTDm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPBROADCASTD")
	}
}

func TestVPBROADCASTQm64byte(t *testing.T) {
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

	VPBROADCASTQm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPBROADCASTQ")
	}
}

func TestVPBROADCASTWm16byte(t *testing.T) {
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

	VPBROADCASTWm16byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPBROADCASTW")
	}
}

func TestVPSLLVDm128byte(t *testing.T) {
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

	VPSLLVDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSLLVD")
	}
}

func TestVPSLLVQm128byte(t *testing.T) {
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

	VPSLLVQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSLLVQ")
	}
}

func TestVPSRAVDm128byte(t *testing.T) {
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

	VPSRAVDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRAVD")
	}
}

func TestVPSRLVDm128byte(t *testing.T) {
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

	VPSRLVDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRLVD")
	}
}

func TestVPSRLVQm128byte(t *testing.T) {
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

	VPSRLVQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRLVQ")
	}
}
