package chapterten

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func TimeRunner() {
	ticker := time.NewTicker(time.Microsecond)
	i := 0

ticker:
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
			i++
			if i == 10 {
				break ticker
			}
		}
	}

	h := handler{}
	fmt.Println(h.getBody1())
}

func (h handler) getStatusCode2(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}

func (h handler) getBody1() (string, error) {
	resp, err := h.client.Get(h.url)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type handler struct {
	client http.Client
	url    string
}

// func handler2(w http.ResponseWriter, req *http.Request) {
//	err := foo(req)
//	if err != nil {
//		http.Error(w, "foo", http.StatusInternalServerError)
//		return
//	}
//
//	_, _ = w.Write([]byte("all good"))
//	w.WriteHeader(http.StatusCreated)
//}
//
//func foo(req *http.Request) error {
//	return nil
//}
