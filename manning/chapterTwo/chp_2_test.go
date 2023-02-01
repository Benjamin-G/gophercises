package chapterTwo

import (
	"bytes"
	"sort"
	"strings"
	"testing"
)

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

func TestCopySourceToDest(t *testing.T) {
	t.Parallel()
	const input = "foo"
	source := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))

	err := copySourceToDest(source, dest)
	if err != nil {
		t.FailNow()
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected: %s, got: %s", input, got)
	}
}

func TestSortOrgans(t *testing.T) {
	t.Parallel()
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	got := isSortedAsc(ByWeight{s})
	//printOrgans(s)
	if got == false {
		t.Errorf("Expected sorted by weight, got:\n %v %v", got, &s)
	}

	sort.Sort(ByName{s})
	got = isSortedAsc(ByName{s})
	//printOrgans(s)
	if got == false {
		t.Errorf("Expected sorted by name, got:\n %v %v", got, &s)
	}

	sort.Slice(s, func(i, j int) bool { return s[i].Weight > s[j].Weight })
	got = isSortedDesc(ByWeight{s})
	//printOrgans(s)
	if got == false {
		t.Errorf("Expected sorted by weight, got:\n %v, %v", got, s)
	}
}
