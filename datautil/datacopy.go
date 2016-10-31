package datautil

import (
	"bytes"
	"encoding/gob"

	"common/pools"
)

/*
bufferPool holds buffers which are used to copy objects.
*/
var bufferPool = pools.NewByteBufferPool()

/*
CopyObject copies contents of a given object reference to another given object reference.
*/
func CopyObject(src interface{}, dest interface{}) error {
	bb := bufferPool.Get().(*bytes.Buffer)

	err := gob.NewEncoder(bb).Encode(src)

	if err != nil {
		return err
	}

	err = gob.NewDecoder(bb).Decode(dest)

	if err != nil {
		return err
	}

	bb.Reset()
	bufferPool.Put(bb)

	return nil
}

/*
MergeMaps merges all given maps into a new map. Contents are shallow copies
and conflicts are resolved as last-one-wins.
*/
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})

	for _, m := range maps {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}
