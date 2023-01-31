package chapterTwo

import (
	"fmt"
	"log"
)

type ErrJoin struct {
	err error
}

func (e ErrJoin) Error() string {
	return fmt.Sprintf("[ErrJoin] %s", e.err)
}

type ErrConcat struct {
	err error
}

func (e ErrConcat) Error() string {
	return fmt.Sprintf("[ErrConcat] %s", e.err)
}

func NestedRunner() {
	s, err := join2("chapter", "hello world", 20)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}

// Code smell
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", ErrJoin{err: fmt.Errorf("s1 is empty")}
	} else {
		if s2 == "" {
			return "", ErrJoin{err: fmt.Errorf("s2 is empty")}
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", ErrJoin{err: fmt.Errorf("s1 is empty")}
	}
	if s2 == "" {
		return "", ErrJoin{err: fmt.Errorf("s2 is empty")}
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", ErrConcat{err: fmt.Errorf("concatenate")}
}
