package main

import "testing"

func BenchmarkCompressBlock(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CompressBlock("Golang is amazing!")
	}
}
