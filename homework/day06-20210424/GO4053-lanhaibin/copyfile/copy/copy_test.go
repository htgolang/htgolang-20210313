package copy

import "testing"

var src string = "/tmp/foo"
var dest string = "/tmp/bar"

func BenchmarkBlock1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(src, dest, 1024)
	}
}

func BenchmarkBlock4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(src, dest, 4096)
	}
}

func BenchmarkBlock8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(src, dest, 8192)
	}
}
func BenchmarkBlock16384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(src, dest, 16384)
	}
}
