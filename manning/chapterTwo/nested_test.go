package chapterTwo

import "testing"

func Test_Join(t *testing.T) {
	t.Parallel()

	// Join 1 Errors
	_, err := join1("", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}
	_, err = join1("hello", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}

	// Join 2 Errors
	_, err = join2("", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}
	_, err = join2("hello", "", 10)
	if _, ok := err.(ErrJoin); !ok {
		t.Fatal("expected ErrJoin, got", err)
	}

	// Concat Errors
	_, err = join2("test", "join", 10)
	if _, ok := err.(ErrConcat); !ok {
		t.Fatal("expected ErrConcat, got", err)
	}
}
