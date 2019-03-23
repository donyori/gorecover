package gorecover

import (
	"errors"
	"testing"
)

func TestRecover1(t *testing.T) {
	t.Log("Test 1")
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
	t.Log("Test 2")
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
	t.Log("Test 3")
	err := Recover(func() {})
	if err != nil {
		t.Errorf("err (%v) != nil", err)
	}
}
