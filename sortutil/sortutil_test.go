package sortutil

import (
	"reflect"
	"testing"
)

func TestInt64s(t *testing.T) {
	testSlice := []int64{5, 2, 3, 0xFFFFFFFF, 1}

	Int64s(testSlice)

	if !reflect.DeepEqual(testSlice, []int64{1, 2, 3, 5, 0xFFFFFFFF}) {
		t.Error("Unexpected sorted order:", testSlice)
	}
}

func TestUInt64s(t *testing.T) {
	testSlice := []uint64{5, 2, 3, 0xFFFFFFFF, 1}

	UInt64s(testSlice)

	if !reflect.DeepEqual(testSlice, []uint64{1, 2, 3, 5, 0xFFFFFFFF}) {
		t.Error("Unexpected sorted order:", testSlice)
	}
}
