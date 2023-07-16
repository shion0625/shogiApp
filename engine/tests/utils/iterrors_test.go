package utils

import (
	"errors"
	"github.com/shion0625/shogiApp/engine/src/utils"
	"testing"
)

func TestWrap(t *testing.T) {
	var err error
	iterrors.Wrap(&err, "whatever")
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	orig := errors.New("bad stuff")
	err = orig
	iterrors.Wrap(&err, "Frob(%d)", 3)
	want := "Frob(3): bad stuff"
	if got := err.Error(); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	if got := errors.Unwrap(err); got != orig {
		t.Errorf("Unwrap: got %#v, want %#v", got, orig)
	}
}
