package fundementals

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func RunInterfaces() {
	m := Musician{Name: "Kurt"}
	PerformAtVenue(m)
	p := Poet{Name: "Janis"}
	PerformAtVenue(p)
}

type Musician struct {
	Name string
}

func (m Musician) Perform() {
	fmt.Println(m.Name, "is singing")
}

type Poet struct {
	Name string
}

func (p Poet) Perform() {
	fmt.Println(p.Name, "is reading poetry")
}

type Performer interface {
	Perform()
}

// shared methods
func PerformAtVenue(m Performer) {
	m.Perform()
}

func WriteDataV1(w *os.File, data []byte) {
	w.Write(data)
}

// Now that WriteData uses the io.Writer, we cannot only use implementations from the
// standard library like os.File and bytes.Buffer, but we can create our own
// implementation of io.Writer.

func WriteDataV2(w io.Writer, data []byte) {
	w.Write(data)
}

type Scribe struct {
	data []byte
}

// By implementing a
// String() string method on the Scribe type, the Scribe now implements both the
// fmt.Stringer and io.Writer interfaces,

func (s Scribe) String() string {
	return string(s.data)
}
func (s *Scribe) Write(p []byte) (int, error) {
	s.data = p
	return len(p), nil
}

var _ io.Writer = &Scribe{}
var _ fmt.Stringer = Scribe{}

func WriteNow(i any) error {
	w, ok := i.(io.Writer)
	if !ok {
		return fmt.Errorf("expected io.Writer, got %T", i)
	}
	now := time.Now()
	w.Write([]byte(now.String()))
	return nil
}

func WriteNowV2(w io.Writer) error {
	now := time.Now()
	// Asserting that io.Writer Is a bytes.Buffer
	if bb, ok := w.(*bytes.Buffer); ok {
		bb.WriteString(now.String())
		return nil
	}
	w.Write([]byte(now.String()))
	return nil
}

// Using a switch Statement to Make Many Type Assertions
func WriteNowV3(i any) error {
	now := time.Now().String()
	switch i.(type) {
	case *bytes.Buffer:
		fmt.Println("type was a *bytes.Buffer", now)
	case io.StringWriter:
		fmt.Println("type was a io.StringWriter", now)
	case io.Writer:
		fmt.Println("type was a io.Writer", now)
	}
	return fmt.Errorf("cannot write to %T", i)
}
