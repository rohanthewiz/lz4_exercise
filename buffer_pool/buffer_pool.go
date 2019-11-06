package buffer_pool

import (
	"bytes"
	"sync"
)

var buffersPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Fetch a new buffer from the pool
func GetBuffer() *bytes.Buffer {
	return buffersPool.Get().(*bytes.Buffer)
}

// Return a buffer to the pool
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffersPool.Put(buf)
}

// Let's check later to see if this works for int[]
var intPool = sync.Pool{
	New: func() interface{} {
		return make([]int, 64<<10)
	},
}

// Fetch a new buffer from the pool
func GetIntArray() []int {
	return intPool.Get().([]int)
}

// Return a buffer to the pool
func PutIntArray(arr []int) {
	//arr = arr[0:0]
	intPool.Put(arr)
}
