package pools

import (
	"bytes"
	"testing"
)

func TestByteBufferPool(t *testing.T) {

	pool := NewByteBufferPool()

	buf1 := pool.Get().(*bytes.Buffer)
	buf2 := pool.Get()
	buf3 := pool.Get()

	if buf1 == nil || buf2 == nil || buf3 == nil {
		t.Error("Initialisation didn't work")
		return
	}

	buf1.Write(make([]byte, 10, 10))

	buf1.Reset()

	pool.Put(buf1)
}
