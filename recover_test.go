package gorecover

import (
	"errors"
	"testing"
)

func TestRecover1(t *testing.T) {
	e := errors.New("error1")
	err := Recover(func() {
		panic(e)
	})
	t.Logf("Error: %v", err)
	if e != err {
		t.Errorf("e (\"%v\" at %p) != err (\"%v\" at %p)", e, e, err, err)
	}
}

func TestRecover2(t *testing.T) {
	e := "error2"
	err := Recover(func() {
		panic(e)
	})
	t.Logf("Error: %v", err)
	if err == nil || err.Error() != e {
		t.Errorf("e (%v) != err (%v)", e, err)
	}
}

func TestRecover3(t *testing.T) {
	err := Recover(func() {})
	if err != nil {
		t.Errorf("err (%v) != nil", err)
	}
}

func TestRecoverWithStackTrace(t *testing.T) {
	e := errors.New("error1")
	err, st := RecoverWithStackTrace(func() {
		panicFunc(e)
	})
	t.Logf("Error: %v", err)
	t.Log("Stack trace:")
	t.Log(st)
	if e != err {
		t.Errorf("e (\"%v\" at %p) != err (\"%v\" at %p)", e, e, err, err)
	}
}

// To test stack trace.
func panicFunc(err interface{}) {
	panic(err)
}
