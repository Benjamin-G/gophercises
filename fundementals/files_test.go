package fundementals

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const dirName = "testdata"

func createTestData(t testing.TB) {
	t.Helper()
	// remove any previous test data
	if err := os.RemoveAll(dirName); err != nil {
		t.Fatal(err)
	}
	// create the data directory
	if err := os.Mkdir(dirName, 0755); err != nil {
		t.Fatal(err)
	}
	list := []string{
		dirName + "/.hidden/d.txt",
		dirName + "/a.txt",
		dirName + "/b.txt",
		dirName + "/e/f/_ignore/i.txt",
		dirName + "/e/f/g.txt",
		dirName + "/e/f/h.txt",
		dirName + "/e/j.txt",
		dirName + "/_testdata/c.txt",
	}
	// create the test data files
	for _, path := range list {
		dir := path
		if ext := filepath.Ext(path); len(ext) > 0 {
			dir = filepath.Dir(path)
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			// ignore if the directory already exists
			if !errors.Is(err, fs.ErrExist) {
				t.Fatal(err)
			}
		}
		fmt.Println("creating", path)
		f, err := os.Create(path)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(f, strings.ToUpper(path))
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}
}

func Test_Walk(t *testing.T) {
	t.Parallel()
	createTestData(t)
	exp := []string{
		dirName + "\\a.txt",
		dirName + "\\b.txt",
		dirName + "\\e\\f\\g.txt",
		dirName + "\\e\\f\\h.txt",
		dirName + "\\e\\j.txt",
	}
	act, err := Walk()
	if err != nil {
		t.Fatal(err)
	}
	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if es != as {
		t.Fatalf("\n[expected] %s\n[got] %s\n", es, as)
	}
}
