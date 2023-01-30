package fundementals

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
)

const dirName = "testdata"

//go:embed c.out
//go:embed testdata
var DataFS embed.FS

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

func createTestFile(t testing.TB, fp string, body []byte) {
	t.Helper()
	// create test data directory
	createTestData(t)
	// assert the file does not exist
	// by trying to stat it.
	// this should return an error
	_, err := os.Stat(fp)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			t.Fatal(err)
		}
	}
	// create the file
	err = Create(fp, body)
	if err != nil {
		t.Fatal(err)
	}
	// read the file into memory
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	act := string(b)
	exp := string(body)
	// assert the file contents are correct
	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func appendTestFile(t testing.TB, fp string, body []byte) {
	t.Helper()
	// read the existing file into memory
	before, err := os.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	// append the new data
	if err := Append(fp, body); err != nil {
		t.Fatal(err)
	}
	// read the new file into memory
	after, err := os.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	// assert the new file contents
	// contain the old and new data
	// Hello, World!Hello, Universe!
	exp := string(append(before, body...))
	act := string(after)
	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func createTestFS(t testing.TB) fstest.MapFS {
	t.Helper()
	cab := fstest.MapFS{}
	files := []string{
		".hidden/d.txt",
		"a.txt",
		"b.txt",
		"e/f/_ignore/i.txt",
		"e/f/g.txt",
		"e/f/h.txt",
		"e/j.txt",
	}
	for _, path := range files {
		cab[path] = &fstest.MapFile{
			Data: []byte(strings.ToUpper(path)),
		}
	}
	return cab
}

func Test_Walk_MockFS(t *testing.T) {
	t.Parallel()
	cab := createTestFS(t)
	exp := []string{
		"a.txt",
		"b.txt",
		"e/f/g.txt",
		"e/f/h.txt",
		"e/j.txt",
	}
	act, err := WalkFS(cab)
	if err != nil {
		t.Fatal(err)
	}
	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if es != as {
		t.Fatalf("\n[expected] %s\n[got] %s\n", es, as)
	}
}

func Test_Walk_Embedded(t *testing.T) {
	t.Parallel()
	createTestData(t)
	exp := []string{
		"c.out",
		dirName + "/a.txt",
		dirName + "/b.txt",
		dirName + "/e/f/g.txt",
		dirName + "/e/f/h.txt",
		dirName + "/e/j.txt",
	}
	act, err := WalkFS(DataFS)
	if err != nil {
		t.Fatal(err)
	}
	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if es != as {
		t.Fatalf("\n[expected] %s\n[got] %s\n", es, as)
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

func Test_WalkFS(t *testing.T) {
	t.Parallel()
	createTestData(t)
	exp := []string{
		"a.txt",
		"b.txt",
		"e/f/g.txt",
		"e/f/h.txt",
		"e/j.txt",
	}
	cab := os.DirFS("testdata")
	act, err := WalkFS(cab)
	if err != nil {
		t.Fatal(err)
	}
	es := strings.Join(exp, ", ")
	as := strings.Join(act, ", ")
	if es != as {
		t.Fatalf("\n[expected] %s\n[got] %s\n", es, as)
	}
}

func Test_Append(t *testing.T) {
	t.Parallel()
	fp := dirName + "/test.txt"
	// create the file and assert
	// the file should now equal the string
	// "Hello, World!"
	createTestFile(t, fp, []byte("Hello, World!"))
	// create the file, again, and assert
	// the file should now equal the string
	// "Hello, Universe!"
	appendTestFile(t, fp, []byte("\nHello, Universe!"))
}

func Test_Read(t *testing.T) {
	t.Parallel()

	fp := dirName + "/test_2.txt"
	// create the file and assert
	// the file should now equal the string
	// "Hello, World!"
	createTestFile(t, fp, []byte("Hello, World!"))

	bb := &bytes.Buffer{}
	err := Read(dirName+"/test_2.txt", bb)
	if err != nil {
		t.Fatal(err)
	}
	exp := "Hello, World!"
	act := bb.String()
	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}
