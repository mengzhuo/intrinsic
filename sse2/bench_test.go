package sse2

import "testing"

func generalMin(a, b []byte) {
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			a[i] = b[i]
		}
	}
}

func genralAnd(a, b []byte) {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] & b[i]
	}
}

func makeP(n int, p bool) (r []byte) {

	r = make([]byte, n)
	if p {
		for i := 0; i < n; i++ {
			r[i] = byte(i % 2)
		}
	} else {
		for i := 0; i < n; i++ {
			r[i] = byte(1 - (i % 2))
		}
	}
	return
}

func BenchmarkPMINUBByte(b *testing.B) {
	a := makeP(16, true)
	c := makeP(16, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PMINUBm128uint8(a, c)
	}
}

func BenchmarkGeneralPMINUBByte(b *testing.B) {
	a := makeP(16, true)
	c := makeP(16, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generalMin(a, c)
	}
}

func BenchmarkPAND(b *testing.B) {
	a := makeP(16, true)
	c := makeP(16, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PANDm128byte(a, c)
	}
}

func BenchmarkGeneralAND(b *testing.B) {
	a := makeP(16, true)
	c := makeP(16, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genralAnd(a, c)
	}
}
