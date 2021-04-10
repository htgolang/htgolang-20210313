package utils

import "testing"

func TestRandText(t *testing.T) {
	r1, r2 := RandText(3), RandText(3)
	if r1 == r2 {
		t.Error("rand text same")
	} else {
		t.Logf("randtext result: %s %s", r1, r2)
	}
}

func BenchmarkInt2StrV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StrV1(i)
	}
}

func BenchmarkInt2StrV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StrV2(i)
	}
}
