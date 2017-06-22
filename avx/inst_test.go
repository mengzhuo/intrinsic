package avx

func TestVADDPDm128byte(t *testing.T) {
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

	VADDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDPD")
	}
}

func TestVADDPDm128float64(t *testing.T) {
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

	VADDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDPD")
	}
}

func TestVADDPSm128byte(t *testing.T) {
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

	VADDPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDPS")
	}
}

func TestVADDPSm128float32(t *testing.T) {
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

	VADDPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDPS")
	}
}

func TestVADDSDm64byte(t *testing.T) {
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

	VADDSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSD")
	}
}

func TestVADDSDm64float64(t *testing.T) {
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

	VADDSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSD")
	}
}

func TestVADDSSm32byte(t *testing.T) {
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

	VADDSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSS")
	}
}

func TestVADDSSm32float32(t *testing.T) {
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

	VADDSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSS")
	}
}

func TestVADDSUBPDm128byte(t *testing.T) {
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

	VADDSUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSUBPD")
	}
}

func TestVADDSUBPDm128float64(t *testing.T) {
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

	VADDSUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSUBPD")
	}
}

func TestVADDSUBPSm128byte(t *testing.T) {
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

	VADDSUBPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSUBPS")
	}
}

func TestVADDSUBPSm128float32(t *testing.T) {
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

	VADDSUBPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VADDSUBPS")
	}
}

func TestVANDNPDm128byte(t *testing.T) {
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

	VANDNPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDNPD")
	}
}

func TestVANDNPDm128float64(t *testing.T) {
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

	VANDNPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDNPD")
	}
}

func TestVANDNPSm128byte(t *testing.T) {
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

	VANDNPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDNPS")
	}
}

func TestVANDNPSm128float32(t *testing.T) {
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

	VANDNPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDNPS")
	}
}

func TestVANDPDm128byte(t *testing.T) {
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

	VANDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDPD")
	}
}

func TestVANDPDm128float64(t *testing.T) {
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

	VANDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDPD")
	}
}

func TestVANDPSm128byte(t *testing.T) {
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

	VANDPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDPS")
	}
}

func TestVANDPSm128float32(t *testing.T) {
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

	VANDPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VANDPS")
	}
}

func TestVCOMISDm64byte(t *testing.T) {
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

	VCOMISDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VCOMISD")
	}
}

func TestVCOMISDm64float64(t *testing.T) {
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

	VCOMISDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VCOMISD")
	}
}

func TestVCOMISSm32byte(t *testing.T) {
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

	VCOMISSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VCOMISS")
	}
}

func TestVCOMISSm32float32(t *testing.T) {
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

	VCOMISSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VCOMISS")
	}
}

func TestVDIVPDm128byte(t *testing.T) {
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

	VDIVPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVPD")
	}
}

func TestVDIVPDm128float64(t *testing.T) {
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

	VDIVPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVPD")
	}
}

func TestVDIVPSm128byte(t *testing.T) {
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

	VDIVPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVPS")
	}
}

func TestVDIVPSm128float32(t *testing.T) {
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

	VDIVPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVPS")
	}
}

func TestVDIVSDm64byte(t *testing.T) {
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

	VDIVSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVSD")
	}
}

func TestVDIVSDm64float64(t *testing.T) {
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

	VDIVSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVSD")
	}
}

func TestVDIVSSm32byte(t *testing.T) {
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

	VDIVSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVSS")
	}
}

func TestVDIVSSm32float32(t *testing.T) {
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

	VDIVSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VDIVSS")
	}
}

func TestVHADDPDm128byte(t *testing.T) {
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

	VHADDPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHADDPD")
	}
}

func TestVHADDPDm128float64(t *testing.T) {
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

	VHADDPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHADDPD")
	}
}

func TestVHADDPSm128byte(t *testing.T) {
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

	VHADDPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHADDPS")
	}
}

func TestVHADDPSm128float32(t *testing.T) {
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

	VHADDPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHADDPS")
	}
}

func TestVHSUBPDm128byte(t *testing.T) {
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

	VHSUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHSUBPD")
	}
}

func TestVHSUBPDm128float64(t *testing.T) {
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

	VHSUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHSUBPD")
	}
}

func TestVHSUBPSm128byte(t *testing.T) {
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

	VHSUBPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHSUBPS")
	}
}

func TestVHSUBPSm128float32(t *testing.T) {
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

	VHSUBPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VHSUBPS")
	}
}

func TestVMASKMOVDQUbyte(t *testing.T) {
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

	VMASKMOVDQUbyte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMASKMOVDQU")
	}
}

func TestVMAXPDm128byte(t *testing.T) {
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

	VMAXPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXPD")
	}
}

func TestVMAXPDm128float64(t *testing.T) {
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

	VMAXPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXPD")
	}
}

func TestVMAXPSm128byte(t *testing.T) {
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

	VMAXPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXPS")
	}
}

func TestVMAXPSm128float32(t *testing.T) {
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

	VMAXPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXPS")
	}
}

func TestVMAXSDm64byte(t *testing.T) {
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

	VMAXSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXSD")
	}
}

func TestVMAXSDm64float64(t *testing.T) {
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

	VMAXSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXSD")
	}
}

func TestVMAXSSm32byte(t *testing.T) {
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

	VMAXSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXSS")
	}
}

func TestVMAXSSm32float32(t *testing.T) {
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

	VMAXSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMAXSS")
	}
}

func TestVMINPDm128byte(t *testing.T) {
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

	VMINPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINPD")
	}
}

func TestVMINPDm128float64(t *testing.T) {
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

	VMINPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINPD")
	}
}

func TestVMINPSm128byte(t *testing.T) {
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

	VMINPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINPS")
	}
}

func TestVMINPSm128float32(t *testing.T) {
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

	VMINPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINPS")
	}
}

func TestVMINSDm64byte(t *testing.T) {
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

	VMINSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINSD")
	}
}

func TestVMINSDm64float64(t *testing.T) {
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

	VMINSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINSD")
	}
}

func TestVMINSSm32byte(t *testing.T) {
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

	VMINSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINSS")
	}
}

func TestVMINSSm32float32(t *testing.T) {
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

	VMINSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMINSS")
	}
}

func TestVMULPDm128byte(t *testing.T) {
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

	VMULPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULPD")
	}
}

func TestVMULPDm128float64(t *testing.T) {
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

	VMULPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULPD")
	}
}

func TestVMULPSm128byte(t *testing.T) {
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

	VMULPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULPS")
	}
}

func TestVMULPSm128float32(t *testing.T) {
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

	VMULPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULPS")
	}
}

func TestVMULSDm64byte(t *testing.T) {
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

	VMULSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULSD")
	}
}

func TestVMULSDm64float64(t *testing.T) {
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

	VMULSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULSD")
	}
}

func TestVMULSSm32byte(t *testing.T) {
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

	VMULSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULSS")
	}
}

func TestVMULSSm32float32(t *testing.T) {
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

	VMULSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VMULSS")
	}
}

func TestVORPDm128byte(t *testing.T) {
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

	VORPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VORPD")
	}
}

func TestVORPDm128float64(t *testing.T) {
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

	VORPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VORPD")
	}
}

func TestVORPSm128byte(t *testing.T) {
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

	VORPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VORPS")
	}
}

func TestVORPSm128float32(t *testing.T) {
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

	VORPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VORPS")
	}
}

func TestVPABSBm128byte(t *testing.T) {
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

	VPABSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSB")
	}
}

func TestVPABSDm128byte(t *testing.T) {
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

	VPABSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSD")
	}
}

func TestVPABSWm128byte(t *testing.T) {
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

	VPABSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPABSW")
	}
}

func TestVPACKSSDWm128byte(t *testing.T) {
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

	VPACKSSDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPACKSSDW")
	}
}

func TestVPACKSSWBm128byte(t *testing.T) {
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

	VPACKSSWBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPACKSSWB")
	}
}

func TestVPACKUSDWm128byte(t *testing.T) {
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

	VPACKUSDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPACKUSDW")
	}
}

func TestVPACKUSWBm128byte(t *testing.T) {
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

	VPACKUSWBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPACKUSWB")
	}
}

func TestVPADDBm128byte(t *testing.T) {
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

	VPADDBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDB")
	}
}

func TestVPADDBm128int8(t *testing.T) {
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

	VPADDBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDB")
	}
}

func TestVPADDDm128byte(t *testing.T) {
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

	VPADDDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDD")
	}
}

func TestVPADDDm128int32(t *testing.T) {
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

	VPADDDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDD")
	}
}

func TestVPADDQm128byte(t *testing.T) {
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

	VPADDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDQ")
	}
}

func TestVPADDQm128int64(t *testing.T) {
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

	VPADDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDQ")
	}
}

func TestVPADDSBm128byte(t *testing.T) {
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

	VPADDSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDSB")
	}
}

func TestVPADDSBm128int8(t *testing.T) {
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

	VPADDSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDSB")
	}
}

func TestVPADDSWm128byte(t *testing.T) {
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

	VPADDSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDSW")
	}
}

func TestVPADDSWm128int16(t *testing.T) {
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

	VPADDSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDSW")
	}
}

func TestVPADDUSBm128byte(t *testing.T) {
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

	VPADDUSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDUSB")
	}
}

func TestVPADDUSBm128uint8(t *testing.T) {
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

	VPADDUSBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDUSB")
	}
}

func TestVPADDUSWm128byte(t *testing.T) {
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

	VPADDUSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDUSW")
	}
}

func TestVPADDUSWm128uint16(t *testing.T) {
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

	VPADDUSWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDUSW")
	}
}

func TestVPADDWm128byte(t *testing.T) {
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

	VPADDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDW")
	}
}

func TestVPADDWm128int16(t *testing.T) {
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

	VPADDWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPADDW")
	}
}

func TestVPANDm128byte(t *testing.T) {
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

	VPANDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPAND")
	}
}

func TestVPANDNm128byte(t *testing.T) {
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

	VPANDNm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPANDN")
	}
}

func TestVPAVGBm128byte(t *testing.T) {
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

	VPAVGBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPAVGB")
	}
}

func TestVPAVGBm128int8(t *testing.T) {
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

	VPAVGBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPAVGB")
	}
}

func TestVPAVGWm128byte(t *testing.T) {
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

	VPAVGWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPAVGW")
	}
}

func TestVPAVGWm128int16(t *testing.T) {
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

	VPAVGWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPAVGW")
	}
}

func TestVPCMPEQBm128byte(t *testing.T) {
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

	VPCMPEQBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPEQB")
	}
}

func TestVPCMPEQDm128byte(t *testing.T) {
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

	VPCMPEQDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPEQD")
	}
}

func TestVPCMPEQQm128byte(t *testing.T) {
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

	VPCMPEQQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPEQQ")
	}
}

func TestVPCMPEQWm128byte(t *testing.T) {
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

	VPCMPEQWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPEQW")
	}
}

func TestVPCMPGTBm128byte(t *testing.T) {
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

	VPCMPGTBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTB")
	}
}

func TestVPCMPGTBm128int8(t *testing.T) {
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

	VPCMPGTBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTB")
	}
}

func TestVPCMPGTDm128byte(t *testing.T) {
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

	VPCMPGTDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTD")
	}
}

func TestVPCMPGTDm128int32(t *testing.T) {
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

	VPCMPGTDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTD")
	}
}

func TestVPCMPGTQm128byte(t *testing.T) {
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

	VPCMPGTQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTQ")
	}
}

func TestVPCMPGTWm128byte(t *testing.T) {
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

	VPCMPGTWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTW")
	}
}

func TestVPCMPGTWm128int16(t *testing.T) {
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

	VPCMPGTWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPCMPGTW")
	}
}

func TestVPERMILPDm128byte(t *testing.T) {
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

	VPERMILPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPERMILPD")
	}
}

func TestVPERMILPDm128float64(t *testing.T) {
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

	VPERMILPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPERMILPD")
	}
}

func TestVPERMILPSm128byte(t *testing.T) {
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

	VPERMILPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPERMILPS")
	}
}

func TestVPERMILPSm128float32(t *testing.T) {
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

	VPERMILPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPERMILPS")
	}
}

func TestVPHADDDm128byte(t *testing.T) {
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

	VPHADDDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHADDD")
	}
}

func TestVPHADDSWm128byte(t *testing.T) {
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

	VPHADDSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHADDSW")
	}
}

func TestVPHADDWm128byte(t *testing.T) {
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

	VPHADDWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHADDW")
	}
}

func TestVPHMINPOSUWm128byte(t *testing.T) {
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

	VPHMINPOSUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHMINPOSUW")
	}
}

func TestVPHSUBDm128byte(t *testing.T) {
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

	VPHSUBDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHSUBD")
	}
}

func TestVPHSUBSWm128byte(t *testing.T) {
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

	VPHSUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHSUBSW")
	}
}

func TestVPHSUBWm128byte(t *testing.T) {
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

	VPHSUBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPHSUBW")
	}
}

func TestVPMADDUBSWm128byte(t *testing.T) {
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

	VPMADDUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMADDUBSW")
	}
}

func TestVPMADDWDm128byte(t *testing.T) {
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

	VPMADDWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMADDWD")
	}
}

func TestVPMADDWDm128int32(t *testing.T) {
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

	VPMADDWDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMADDWD")
	}
}

func TestVPMAXSBm128byte(t *testing.T) {
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

	VPMAXSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSB")
	}
}

func TestVPMAXSBm128int8(t *testing.T) {
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

	VPMAXSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSB")
	}
}

func TestVPMAXSDm128byte(t *testing.T) {
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

	VPMAXSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSD")
	}
}

func TestVPMAXSDm128int32(t *testing.T) {
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

	VPMAXSDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSD")
	}
}

func TestVPMAXSWm128byte(t *testing.T) {
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

	VPMAXSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSW")
	}
}

func TestVPMAXSWm128int16(t *testing.T) {
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

	VPMAXSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXSW")
	}
}

func TestVPMAXUBm128byte(t *testing.T) {
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

	VPMAXUBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUB")
	}
}

func TestVPMAXUBm128uint8(t *testing.T) {
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

	VPMAXUBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUB")
	}
}

func TestVPMAXUDm128byte(t *testing.T) {
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

	VPMAXUDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUD")
	}
}

func TestVPMAXUDm128uint32(t *testing.T) {
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

	VPMAXUDm128uint32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUD")
	}
}

func TestVPMAXUWm128byte(t *testing.T) {
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

	VPMAXUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUW")
	}
}

func TestVPMAXUWm128uint16(t *testing.T) {
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

	VPMAXUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMAXUW")
	}
}

func TestVPMINSBm128byte(t *testing.T) {
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

	VPMINSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSB")
	}
}

func TestVPMINSBm128int8(t *testing.T) {
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

	VPMINSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSB")
	}
}

func TestVPMINSDm128byte(t *testing.T) {
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

	VPMINSDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSD")
	}
}

func TestVPMINSDm128int32(t *testing.T) {
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

	VPMINSDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSD")
	}
}

func TestVPMINSWm128byte(t *testing.T) {
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

	VPMINSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSW")
	}
}

func TestVPMINSWm128int16(t *testing.T) {
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

	VPMINSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINSW")
	}
}

func TestVPMINUBm128byte(t *testing.T) {
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

	VPMINUBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUB")
	}
}

func TestVPMINUBm128uint8(t *testing.T) {
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

	VPMINUBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUB")
	}
}

func TestVPMINUDm128byte(t *testing.T) {
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

	VPMINUDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUD")
	}
}

func TestVPMINUDm128uint32(t *testing.T) {
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

	VPMINUDm128uint32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUD")
	}
}

func TestVPMINUWm128byte(t *testing.T) {
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

	VPMINUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUW")
	}
}

func TestVPMINUWm128uint16(t *testing.T) {
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

	VPMINUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMINUW")
	}
}

func TestVPMOVSXBDm32byte(t *testing.T) {
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

	VPMOVSXBDm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXBD")
	}
}

func TestVPMOVSXBQm16byte(t *testing.T) {
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

	VPMOVSXBQm16byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXBQ")
	}
}

func TestVPMOVSXBWm64byte(t *testing.T) {
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

	VPMOVSXBWm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXBW")
	}
}

func TestVPMOVSXDQm64byte(t *testing.T) {
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

	VPMOVSXDQm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXDQ")
	}
}

func TestVPMOVSXWDm64byte(t *testing.T) {
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

	VPMOVSXWDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXWD")
	}
}

func TestVPMOVSXWQm32byte(t *testing.T) {
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

	VPMOVSXWQm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVSXWQ")
	}
}

func TestVPMOVZXBDm32byte(t *testing.T) {
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

	VPMOVZXBDm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXBD")
	}
}

func TestVPMOVZXBQm16byte(t *testing.T) {
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

	VPMOVZXBQm16byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXBQ")
	}
}

func TestVPMOVZXBWm64byte(t *testing.T) {
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

	VPMOVZXBWm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXBW")
	}
}

func TestVPMOVZXDQm64byte(t *testing.T) {
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

	VPMOVZXDQm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXDQ")
	}
}

func TestVPMOVZXWDm64byte(t *testing.T) {
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

	VPMOVZXWDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXWD")
	}
}

func TestVPMOVZXWQm32byte(t *testing.T) {
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

	VPMOVZXWQm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMOVZXWQ")
	}
}

func TestVPMULDQm128byte(t *testing.T) {
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

	VPMULDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULDQ")
	}
}

func TestVPMULDQm128int64(t *testing.T) {
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

	VPMULDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULDQ")
	}
}

func TestVPMULHRSWm128byte(t *testing.T) {
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

	VPMULHRSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULHRSW")
	}
}

func TestVPMULHUWm128byte(t *testing.T) {
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

	VPMULHUWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULHUW")
	}
}

func TestVPMULHUWm128uint16(t *testing.T) {
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

	VPMULHUWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULHUW")
	}
}

func TestVPMULHWm128byte(t *testing.T) {
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

	VPMULHWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULHW")
	}
}

func TestVPMULHWm128int16(t *testing.T) {
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

	VPMULHWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULHW")
	}
}

func TestVPMULLDm128byte(t *testing.T) {
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

	VPMULLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULLD")
	}
}

func TestVPMULLDm128int32(t *testing.T) {
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

	VPMULLDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULLD")
	}
}

func TestVPMULLWm128byte(t *testing.T) {
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

	VPMULLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULLW")
	}
}

func TestVPMULLWm128int16(t *testing.T) {
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

	VPMULLWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULLW")
	}
}

func TestVPMULUDQm128byte(t *testing.T) {
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

	VPMULUDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULUDQ")
	}
}

func TestVPMULUDQm128int64(t *testing.T) {
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

	VPMULUDQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPMULUDQ")
	}
}

func TestVPORm128byte(t *testing.T) {
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

	VPORm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPOR")
	}
}

func TestVPSADBWm128byte(t *testing.T) {
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

	VPSADBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSADBW")
	}
}

func TestVPSHUFBm128byte(t *testing.T) {
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

	VPSHUFBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSHUFB")
	}
}

func TestVPSIGNBm128byte(t *testing.T) {
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

	VPSIGNBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSIGNB")
	}
}

func TestVPSIGNDm128byte(t *testing.T) {
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

	VPSIGNDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSIGND")
	}
}

func TestVPSIGNWm128byte(t *testing.T) {
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

	VPSIGNWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSIGNW")
	}
}

func TestVPSLLDm128byte(t *testing.T) {
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

	VPSLLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSLLD")
	}
}

func TestVPSLLQm128byte(t *testing.T) {
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

	VPSLLQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSLLQ")
	}
}

func TestVPSLLWm128byte(t *testing.T) {
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

	VPSLLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSLLW")
	}
}

func TestVPSRADm128byte(t *testing.T) {
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

	VPSRADm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRAD")
	}
}

func TestVPSRAWm128byte(t *testing.T) {
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

	VPSRAWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRAW")
	}
}

func TestVPSRLDm128byte(t *testing.T) {
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

	VPSRLDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRLD")
	}
}

func TestVPSRLQm128byte(t *testing.T) {
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

	VPSRLQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRLQ")
	}
}

func TestVPSRLWm128byte(t *testing.T) {
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

	VPSRLWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSRLW")
	}
}

func TestVPSUBBm128byte(t *testing.T) {
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

	VPSUBBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBB")
	}
}

func TestVPSUBBm128int8(t *testing.T) {
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

	VPSUBBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBB")
	}
}

func TestVPSUBDm128byte(t *testing.T) {
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

	VPSUBDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBD")
	}
}

func TestVPSUBDm128int32(t *testing.T) {
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

	VPSUBDm128int32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBD")
	}
}

func TestVPSUBQm128byte(t *testing.T) {
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

	VPSUBQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBQ")
	}
}

func TestVPSUBQm128int64(t *testing.T) {
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

	VPSUBQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBQ")
	}
}

func TestVPSUBSBm128byte(t *testing.T) {
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

	VPSUBSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBSB")
	}
}

func TestVPSUBSBm128int8(t *testing.T) {
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

	VPSUBSBm128int8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBSB")
	}
}

func TestVPSUBSWm128byte(t *testing.T) {
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

	VPSUBSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBSW")
	}
}

func TestVPSUBSWm128int16(t *testing.T) {
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

	VPSUBSWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBSW")
	}
}

func TestVPSUBUSBm128byte(t *testing.T) {
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

	VPSUBUSBm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBUSB")
	}
}

func TestVPSUBUSBm128uint8(t *testing.T) {
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

	VPSUBUSBm128uint8(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBUSB")
	}
}

func TestVPSUBUSWm128byte(t *testing.T) {
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

	VPSUBUSWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBUSW")
	}
}

func TestVPSUBUSWm128uint16(t *testing.T) {
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

	VPSUBUSWm128uint16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBUSW")
	}
}

func TestVPSUBWm128byte(t *testing.T) {
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

	VPSUBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBW")
	}
}

func TestVPSUBWm128int16(t *testing.T) {
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

	VPSUBWm128int16(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPSUBW")
	}
}

func TestVPTESTm128byte(t *testing.T) {
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

	VPTESTm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPTEST")
	}
}

func TestVPTESTm256byte(t *testing.T) {
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

	VPTESTm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPTEST")
	}
}

func TestVPUNPCKHBWm128byte(t *testing.T) {
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

	VPUNPCKHBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKHBW")
	}
}

func TestVPUNPCKHDQm128byte(t *testing.T) {
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

	VPUNPCKHDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKHDQ")
	}
}

func TestVPUNPCKHQDQm128byte(t *testing.T) {
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

	VPUNPCKHQDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKHQDQ")
	}
}

func TestVPUNPCKHWDm128byte(t *testing.T) {
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

	VPUNPCKHWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKHWD")
	}
}

func TestVPUNPCKLBWm128byte(t *testing.T) {
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

	VPUNPCKLBWm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKLBW")
	}
}

func TestVPUNPCKLDQm128byte(t *testing.T) {
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

	VPUNPCKLDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKLDQ")
	}
}

func TestVPUNPCKLQDQm128byte(t *testing.T) {
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

	VPUNPCKLQDQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKLQDQ")
	}
}

func TestVPUNPCKLWDm128byte(t *testing.T) {
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

	VPUNPCKLWDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPUNPCKLWD")
	}
}

func TestVPXORm128byte(t *testing.T) {
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

	VPXORm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VPXOR")
	}
}

func TestVRCPPSm128byte(t *testing.T) {
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

	VRCPPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPPS")
	}
}

func TestVRCPPSm128float32(t *testing.T) {
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

	VRCPPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPPS")
	}
}

func TestVRCPPSm256byte(t *testing.T) {
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

	VRCPPSm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPPS")
	}
}

func TestVRCPPSm256float32(t *testing.T) {
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

	VRCPPSm256float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPPS")
	}
}

func TestVRCPSSm32byte(t *testing.T) {
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

	VRCPSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPSS")
	}
}

func TestVRCPSSm32float32(t *testing.T) {
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

	VRCPSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRCPSS")
	}
}

func TestVRSQRTPSm128byte(t *testing.T) {
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

	VRSQRTPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTPS")
	}
}

func TestVRSQRTPSm128float32(t *testing.T) {
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

	VRSQRTPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTPS")
	}
}

func TestVRSQRTPSm256byte(t *testing.T) {
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

	VRSQRTPSm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTPS")
	}
}

func TestVRSQRTPSm256float32(t *testing.T) {
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

	VRSQRTPSm256float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTPS")
	}
}

func TestVRSQRTSSm32byte(t *testing.T) {
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

	VRSQRTSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTSS")
	}
}

func TestVRSQRTSSm32float32(t *testing.T) {
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

	VRSQRTSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VRSQRTSS")
	}
}

func TestVSQRTPDm128byte(t *testing.T) {
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

	VSQRTPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPD")
	}
}

func TestVSQRTPDm128float64(t *testing.T) {
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

	VSQRTPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPD")
	}
}

func TestVSQRTPDm256byte(t *testing.T) {
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

	VSQRTPDm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPD")
	}
}

func TestVSQRTPDm256float64(t *testing.T) {
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

	VSQRTPDm256float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPD")
	}
}

func TestVSQRTPSm128byte(t *testing.T) {
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

	VSQRTPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPS")
	}
}

func TestVSQRTPSm128float32(t *testing.T) {
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

	VSQRTPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPS")
	}
}

func TestVSQRTPSm256byte(t *testing.T) {
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

	VSQRTPSm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPS")
	}
}

func TestVSQRTPSm256float32(t *testing.T) {
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

	VSQRTPSm256float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTPS")
	}
}

func TestVSQRTSDm64byte(t *testing.T) {
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

	VSQRTSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTSD")
	}
}

func TestVSQRTSDm64float64(t *testing.T) {
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

	VSQRTSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTSD")
	}
}

func TestVSQRTSSm32byte(t *testing.T) {
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

	VSQRTSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSQRTSS")
	}
}

func TestVSUBPDm128byte(t *testing.T) {
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

	VSUBPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBPD")
	}
}

func TestVSUBPDm128float64(t *testing.T) {
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

	VSUBPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBPD")
	}
}

func TestVSUBPSm128byte(t *testing.T) {
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

	VSUBPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBPS")
	}
}

func TestVSUBPSm128float32(t *testing.T) {
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

	VSUBPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBPS")
	}
}

func TestVSUBSDm64byte(t *testing.T) {
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

	VSUBSDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBSD")
	}
}

func TestVSUBSDm64float64(t *testing.T) {
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

	VSUBSDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBSD")
	}
}

func TestVSUBSSm32byte(t *testing.T) {
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

	VSUBSSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBSS")
	}
}

func TestVSUBSSm32float32(t *testing.T) {
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

	VSUBSSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VSUBSS")
	}
}

func TestVTESTPDm128byte(t *testing.T) {
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

	VTESTPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VTESTPD")
	}
}

func TestVTESTPDm256byte(t *testing.T) {
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

	VTESTPDm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VTESTPD")
	}
}

func TestVTESTPSm256byte(t *testing.T) {
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

	VTESTPSm256byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VTESTPS")
	}
}

func TestVTESTPSm128byte(t *testing.T) {
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

	VTESTPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VTESTPS")
	}
}

func TestVUCOMISDm64byte(t *testing.T) {
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

	VUCOMISDm64byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUCOMISD")
	}
}

func TestVUCOMISDm64float64(t *testing.T) {
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

	VUCOMISDm64float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUCOMISD")
	}
}

func TestVUCOMISSm32byte(t *testing.T) {
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

	VUCOMISSm32byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUCOMISS")
	}
}

func TestVUCOMISSm32float32(t *testing.T) {
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

	VUCOMISSm32float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUCOMISS")
	}
}

func TestVUNPCKHPDm128byte(t *testing.T) {
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

	VUNPCKHPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKHPD")
	}
}

func TestVUNPCKHPDm128float64(t *testing.T) {
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

	VUNPCKHPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKHPD")
	}
}

func TestVUNPCKHPSm128byte(t *testing.T) {
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

	VUNPCKHPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKHPS")
	}
}

func TestVUNPCKHPSm128float32(t *testing.T) {
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

	VUNPCKHPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKHPS")
	}
}

func TestVUNPCKLPDm128byte(t *testing.T) {
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

	VUNPCKLPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKLPD")
	}
}

func TestVUNPCKLPDm128float64(t *testing.T) {
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

	VUNPCKLPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKLPD")
	}
}

func TestVUNPCKLPSm128byte(t *testing.T) {
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

	VUNPCKLPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKLPS")
	}
}

func TestVUNPCKLPSm128float32(t *testing.T) {
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

	VUNPCKLPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VUNPCKLPS")
	}
}

func TestVXORPDm128byte(t *testing.T) {
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

	VXORPDm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VXORPD")
	}
}

func TestVXORPDm128float64(t *testing.T) {
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

	VXORPDm128float64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VXORPD")
	}
}

func TestVXORPSm128byte(t *testing.T) {
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

	VXORPSm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VXORPS")
	}
}

func TestVXORPSm128float32(t *testing.T) {
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

	VXORPSm128float32(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on VXORPS")
	}
}
