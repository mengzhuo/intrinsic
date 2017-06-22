package sse42

func TestPCMPGTQm128byte(t *testing.T) {
	a := make([]byte, 64)
	aT := make([]byte, 64)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	copy(aT, a)

	b := make([]byte, 64)
	bT := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 2
	}
	copy(bT, b)

	PCMPGTQm128byte(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on PCMPGTQ")
	}
}

func TestPCMPGTQm128int64(t *testing.T) {
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

	PCMPGTQm128int64(a, b)
	if a[0] == aT[0] && b[0] == bT[0] {
		t.Error("Nothing changed on PCMPGTQ")
	}
}
