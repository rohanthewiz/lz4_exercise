package main

import (
	"log"
	"lz4_exercise/buffer_pool"
	"strings"
	"time"

	"github.com/pierrec/lz4"
)

func main() {
	println("LZ4 Compression test")
	CompressBlock("hello world")
	CompressBlock("Golang is awesome!")
	CompressBlock("God is good!")
}

func CompressBlock(s string) {
	data := []byte(strings.Repeat(s, 50))
	//log.Println("data length is", len(data))

	// COMPRESSION
	t0 := time.Now()
	compBuf := buffer_pool.GetBuffer()
	defer buffer_pool.PutBuffer(compBuf) // remember to put it back
	compBuf.Grow(len(data))
	buf := compBuf.Bytes()

	// Buffer pool for hash table
	ht := buffer_pool.GetIntArray()
	defer buffer_pool.PutIntArray(ht)
	//ht := make([]int, 64<<10) // buffer for the (hash) table

	n, err := lz4.CompressBlock(data, buf[0:len(data)], ht)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Compression took", time.Since(t0))

	buf = buf[0:n]
	if n == 0 || n >= len(data) {
		log.Printf("`%s` is not compressible", s)
		return
	} else {
		log.Println("Compressed length is", n, " ratio:", float64(n)/float64(len(data)))
	}

	// DECOMPRESSION
	// TODO - Use buffer pool for decompression
	t0 = time.Now()
	// Allocated a very large buffer for decompression.
	out := make([]byte, 10*len(data))
	n, err = lz4.UncompressBlock(buf, out)
	if err != nil {
		log.Println(err)
	}
	out = out[:n]
	log.Println("Decompression took", time.Since(t0))

	// Output:
	// hello world
	log.Println(string(out[:len(s)]))
}
