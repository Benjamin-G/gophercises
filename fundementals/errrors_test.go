package fundementals

import "testing"

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
