package errorutil

import (
	"errors"
	"testing"
)

func TestCompositeError(t *testing.T) {

	ce := NewCompositeError()

	if ce.HasErrors() {
		t.Error("CompositeError object shouldn't have any errors yet")
		return
	}

	ce.Add(errors.New("test1"))

	if !ce.HasErrors() {
		t.Error("CompositeError object should have one error by now")
		return
	}

	ce.Add(errors.New("test2"))

	// Add a CompositeError to a CompositeError

	ce2 := NewCompositeError()
	ce2.Add(errors.New("test3"))
	ce.Add(ce2)

	if ce.Error() != "test1; test2; test3" {
		t.Error("Unexpected output:", ce.Error())
	}
}
