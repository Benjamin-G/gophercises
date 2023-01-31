package chapterTwo

import "testing"

func Test_Join(t *testing.T) {
	t.Parallel()
	_, err := join1("", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}
	_, err = join1("hello", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}
}
