package chaptereleven

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// func TestLongRunning(t *testing.T) {
//	if testing.Short() {
//		t.Skip("skipping long-running test")
//	}
//	// ...
//}

func TestRemoveNewLineSuffix(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		`empty`: {
			input:    "",
			expected: "",
		},
		`ending with \r\n`: {
			input:    "a\r\n",
			expected: "a",
		},
		`ending with \n`: {
			input:    "a\n",
			expected: "a",
		},
		`ending with multiple \n`: {
			input:    "a\n\n\n",
			expected: "a",
		},
		`ending without newline`: {
			input:    "a",
			expected: "a",
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := removeNewLineSuffixes(tt.input)
			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}

// BAD
type publisherMock1 struct {
	mu  sync.RWMutex
	got []Foo
}

func (p *publisherMock1) Publish(got []Foo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = got
}

func (p *publisherMock1) Get() []Foo {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.got
}

func TestGetBestFoo(t *testing.T) {
	mock := publisherMock1{}
	h := Handler{
		publisher: &mock,
		n:         2,
	}

	foo := h.getBestFoo(42)
	// Check foo
	_ = foo

	time.Sleep(10 * time.Millisecond)

	published := mock.Get()
	// Check published
	_ = published

	assert(t, func() bool {
		return len(mock.Get()) == 2
	}, 30, time.Millisecond)
}

func assert(t *testing.T, assertion func() bool,
	maxRetry int, waitTime time.Duration) {
	t.Helper()

	for i := 0; i < maxRetry; i++ {
		if assertion() {
			return
		}

		time.Sleep(waitTime)
	}
	t.Fail()
}

// GOOD
type publisherMock2 struct {
	ch chan []Foo
}

func (p *publisherMock2) Publish(got []Foo) {
	p.ch <- got
}

func TestGetBestFoo2(t *testing.T) {
	mock := publisherMock2{
		ch: make(chan []Foo),
	}
	defer close(mock.ch)

	h := Handler{
		publisher: &mock,
		n:         2,
	}
	foo := h.getBestFoo(42)
	// Check foo
	_ = foo

	if v := len(<-mock.ch); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
}

func TestCache_TrimOlderThan(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.04Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.05Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.06Z")},
	}
	cache := &Cache{}
	cache.Add(events)
	cache.TrimOlderThan(parseTime(t, "2020-01-01T12:00:00.06Z").
		Add(-15 * time.Millisecond))

	got := cache.GetAll()

	expected := 2

	if len(got) != expected {
		t.Fatalf("expected %d, got %d", expected, len(got))
	}
}

func parseTime(t *testing.T, timestamp string) time.Time {
	t.Helper()

	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t.FailNow()
	}

	return ts
}

func TestHandler2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost",
		strings.NewReader("foo"))
	w := httptest.NewRecorder()
	Handler2(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	body, _ := io.ReadAll(w.Body)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if http.StatusOK != w.Result().StatusCode {
		t.FailNow()
	}
}

func TestDurationClientGet(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"duration": 314}`))
			},
		),
	)
	defer srv.Close()

	client := NewDurationClient()
	duration, err :=
		client.GetDuration(srv.URL, 51.551261, -0.1221146, 51.57, -0.13)

	if err != nil {
		t.Fatal(err)
	}

	if duration != 314*time.Second {
		t.Errorf("expected 314 seconds, got %v", duration)
	}
}
