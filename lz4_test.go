package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
func BenchmarkCompressBlock(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CompressBlock("Golang is amazing!")
		CompressBlock("Jesus, son of God")
	}
}
