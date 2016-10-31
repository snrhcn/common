package pools

import (
	"bytes"
	"sync"
)

/*
NewByteBufferPool creates a new pool of bytes.Buffer objects. The pool creates
new ones if it runs empty.
*/
func NewByteBufferPool() *sync.Pool {
	return &sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}
}
