package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/pierrec/lz4"
)

func main() {
	println("LZ4 Compression test")
	compressBlock("hello world")
}

func compressBlock(s string) {
	data := []byte(strings.Repeat(s, 50))
	fmt.Println("data length is", len(data))

	// COMPRESSION
	t0 := time.Now()
	compBuf := bytes.NewBuffer(make([]byte, len(data))) // for compressed output

	ht := make([]int, 64<<10) // buffer for the (hash) table

	n, err := lz4.CompressBlock(data, compBuf.Bytes(), ht)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compression took", time.Since(t0))

	compBuf.Truncate(n)
	if n == 0 || n >= len(data) {
		fmt.Printf("`%s` is not compressible", s)
		return
	} else {
		fmt.Println("Compressed length is", n, " ratio:", float64(n) / float64(len(data)))
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
