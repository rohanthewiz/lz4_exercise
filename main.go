package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/pierrec/lz4"
)

func main() {
	println("LZ4 Compression test")
	b := Compress("hello world,hello world,hello world")
	fmt.Println(Decompress(b))
	b = Compress("God is good! God is good! God is good! God is good! ")
	fmt.Println(Decompress(b))
}

func Compress(s string) (cBuf bytes.Buffer) {
	t0 := time.Now()
	rdr := bytes.NewReader([]byte(s))

	//cBuf = buffer_pool.GetBuffer()
	zwriter := lz4.NewWriter(&cBuf)
	defer func() {
		err := zwriter.Close()
		if err != nil {
			log.Println("error closing compress writer", err.Error())
			return
		}
	}()
	_, err := io.Copy(zwriter, rdr)
	if err != nil {
		log.Println("error on compress", err.Error())
		return
	}

	log.Println("Compression took", time.Since(t0))

	if cBuf.Len() >= len([]byte(s)) { // TODO - Optimz
		log.Printf("`%s` is not compressible", s)
		return
	}
	log.Println("Compressed length is", cBuf.Len(), " ratio:", float64(cBuf.Len())/float64(len([]byte(s))))
	return
}

func Decompress(inBuf bytes.Buffer) (s string) {
	//defer buffer_pool.PutBuffer(inBuf) // remember to put it back
	t0 := time.Now()

	var outBuf bytes.Buffer
	//outBuf := buffer_pool.GetBuffer()
	//defer buffer_pool.PutBuffer(outBuf)

	zr := lz4.NewReader(&inBuf)
	_, err := io.Copy(&outBuf, zr)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Decompression took", time.Since(t0))

	return string(outBuf.Bytes())
}
