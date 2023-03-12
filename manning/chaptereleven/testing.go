package chaptereleven

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

func removeNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}

	if strings.HasSuffix(s, "\r\n") {
		return removeNewLineSuffixes(s[:len(s)-2])
	}

	if strings.HasSuffix(s, "\n") {
		return removeNewLineSuffixes(s[:len(s)-1])
	}

	return s
}

type Handler struct {
	n         int
	publisher publisher
}

type publisher interface {
	Publish([]Foo)
}

func (h Handler) getBestFoo(someInputs int) Foo {
	foos := getFoos(someInputs)
	best := foos[0]

	go func() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}

		h.publisher.Publish(foos)
	}()

	return best
}

func getFoos(inputs int) []Foo {
	return make([]Foo, inputs)
}

type Foo struct{}

type Cache struct {
	mu     sync.RWMutex
	events []Event
}

type Event struct {
	Timestamp time.Time
	Data      string
}

func (c *Cache) TrimOlderThan(t time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

func (c *Cache) Add(events []Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events = append(c.events, events...)
}

func (c *Cache) GetAll() []Event {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.events
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func (c DurationClient) GetDuration(url string, lat1, lng1, lat2, lng2 float64) (time.Duration, error) {
	resp, err := c.client.Post(
		url, "application/json",
		buildRequestBody(lat1, lng1, lat2, lng2),
	)
	if err != nil {
		return 0, err
	}

	return parseResponseBody(resp.Body)
}

type request struct {
	Duration int
}

func buildRequestBody(lat1, lng1, lat2, lng2 float64) io.Reader {
	return strings.NewReader("")
}

type DurationClient struct {
	client *http.Client
}

func NewDurationClient() DurationClient {
	return DurationClient{
		client: http.DefaultClient,
	}
}

func parseResponseBody(r io.ReadCloser) (time.Duration, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = r.Close()
	}()

	var req request
	err = json.Unmarshal(b, &req)
	if err != nil {
		return 0, err
	}
	return time.Duration(req.Duration) * time.Second, nil
}

type LowerCaseReader struct {
	reader io.Reader
}

func (l LowerCaseReader) Read(p []byte) (int, error) {
	return 0, nil
}

func foo1(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// ...
	_ = b
	return nil
}

func foo2(r io.Reader) error {
	b, err := readAll(r, 3)
	if err != nil {
		return err
	}

	// ...
	_ = b
	return nil
}

func readAll(r io.Reader, retries int) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			retries--
			if retries < 0 {
				return b, err
			}
		}
	}
}

type Customer struct {
	id string
}
