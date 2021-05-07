package main

import (
	"fmt"
	"testing"
)

func BenchmarkCopyFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
}
