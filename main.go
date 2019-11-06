package main

import (
	"fmt"
	"lz4_exercise/buffer_pool"
	"strings"
	"time"

	"github.com/pierrec/lz4"
)

func main() {
	println("LZ4 Compression test")
	CompressBlock("hello world")
	CompressBlock("Golang is awesome!")
}

func CompressBlock(s string) {
	data := []byte(strings.Repeat(s, 50))
	fmt.Println("data length is", len(data))

	// COMPRESSION
	t0 := time.Now()
	compBuf := buffer_pool.GetBuffer()
	defer buffer_pool.PutBuffer(compBuf) // remember to put it back
	compBuf.Grow(len(data))
	buf := compBuf.Bytes()

	ht := make([]int, 64<<10) // buffer for the (hash) table

	n, err := lz4.CompressBlock(data, buf[0:len(data)], ht)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compression took", time.Since(t0))

	buf = buf[0:n]
	if n == 0 || n >= len(data) {
		fmt.Printf("`%s` is not compressible", s)
		return
	} else {
		fmt.Println("Compressed length is", n, " ratio:", float64(n)/float64(len(data)))
	}

	// DECOMPRESSION
	t0 = time.Now()
	// Allocated a very large buffer for decompression.
	out := make([]byte, 10*len(data))
	n, err = lz4.UncompressBlock(compBuf.Bytes(), out)
	if err != nil {
		fmt.Println(err)
	}
	out = out[:n]
	fmt.Println("Decompression took", time.Since(t0))

	// Output:
	// hello world
	fmt.Println(string(out[:len(s)]))
}
