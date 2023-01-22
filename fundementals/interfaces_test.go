package fundementals

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// we need to create a new
// file, call the WriteData function, close the file, reopen the file, read the file, and then
// compare the contents.
func Test_WriteDataV1(t *testing.T) {
	t.Parallel()
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal(err)
	}
	fn := filepath.Join(dir, "hello.txt")
	f, err := os.Create(fn)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("Hello, World!")
	WriteDataV1(f, data)
	f.Close()
	f, err = os.Open(fn)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	act := string(b)
	exp := string(data)
	if act != exp {
		t.Fatalf("expected %q, got %q", exp, act)
	}
}

func Test_WriteDataV2(t *testing.T) {
	t.Parallel()
	// create a buffer to write to
	bb := &bytes.Buffer{}
	data := []byte("Hello, World!")
	// write the data to the buffer
	WriteDataV2(bb, data)
	// capture the data written to the buffer
	// to the act variable
	act := bb.String()
	exp := string(data)
	// compare the expected and actual values
	if act != exp {
		t.Fatalf("expected %q, got %q", exp, act)
	}
}

func Test_ScribeWriteData(t *testing.T) {
	t.Parallel()
	scribe := &Scribe{}
	data := []byte("Hello, World!")
	WriteDataV2(scribe, data)
	act := scribe.String()
	exp := string(data)
	if act != exp {
		t.Fatalf("expected %q, got %q", exp, act)
	}
}
