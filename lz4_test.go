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
		Compress("Golang is amazing!Golang is amazing!Golang is amazing!Golang is amazing!Golang is amazing!")
		Compress("Golang is amazing!Golang is amazing!Golang is amazing!Golang is amazing!Golang is amazing!")
		Compress("Jesus, son of GodJesus, son of GodJesus, son of GodJesus, son of GodJesus, son of GodJesus, son of God")
	}
}
