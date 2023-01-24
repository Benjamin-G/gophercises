package fundementals

import (
	"errors"
	"testing"
)

func Test_Get(t *testing.T) {
	t.Parallel()
	act, err := GetString("a")
	if err != nil {
		t.Fatalf("expect no error, got %s", err)
	}
	exp := "A"
	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
	_, err = GetString("?")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func Test_DoSomething(t *testing.T) {
	t.Parallel()
	err := doSomething(0)
	if err != nil {
		t.Fatal(err)
	}

	// Proper testing
	err = doSomething(1)
	if _, ok := err.(ErrNonErrCaught); !ok {
		t.Fatal("expected ErrNonErrCaught, got", err)
	}

	err = doSomething(2)
	if err == nil {
		t.Fatal("expected err, got nil")
	}
}

func Test_Admin(t *testing.T) {
	t.Parallel()
	a := &Admin{
		User: &User{name: "Kurt"},
	}
	if a.String() != "Kurt" {
		t.Fatal("expected Kurt")
	}
	b := &Admin{}
	if b.User != nil {
		t.Fatal("expected to not have a User")
	}
}

func Test_Unwrap(t *testing.T) {
	t.Parallel()
	original := errors.New("original error")
	wrapped := Wrapper(original)
	unwrapped := errors.Unwrap(wrapped)
	if unwrapped != original {
		t.Fatalf("expected %v, got %v", original, unwrapped)
	}
}

func Test_As(t *testing.T) {
	t.Parallel()
	original := errors.New("original error")
	wrapped := Wrapper(original)
	actA := ErrorA{}
	ok := errors.As(wrapped, &actA)
	if !ok {
		t.Fatalf("expected %v to act as %v", wrapped, actA)
	}
	if actA.err == nil {
		t.Fatalf("expected non-nil, got nil")
	}

	actB := ErrorB{}
	ok = errors.As(wrapped, &actB)
	if !ok {
		t.Fatalf("expected %v to act as %v", wrapped, actB)
	}
	if actB.err == nil {
		t.Fatalf("expected non-nil, got nil")
	}

	actC := ErrorC{}
	ok = errors.As(wrapped, &actC)
	if !ok {
		t.Fatalf("expected %v to act as %v", wrapped, actC)
	}
	if actC.err == nil {
		t.Fatalf("expected non-nil, got nil")
	}
}

func Test_Is(t *testing.T) {
	t.Parallel()
	original := errors.New("original error")
	wrapped := Wrapper(original)
	expA := ErrorA{}
	ok := errors.Is(wrapped, expA)
	if !ok {
		t.Fatalf("expected %v to be %v", wrapped, expA)
	}

	expB := ErrorB{}
	ok = errors.Is(wrapped, expB)
	if !ok {
		t.Fatalf("expected %v to be %v", wrapped, expB)
	}

	expC := ErrorC{}
	ok = errors.Is(wrapped, expC)
	if !ok {
		t.Fatalf("expected %v to be %v", wrapped, expC)
	}
}
