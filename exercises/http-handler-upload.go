package exercises

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// https://drstearns.github.io/tutorials/gomiddleware/
// https://tutorialedge.net/golang/go-file-upload-tutorial/

// HTTP handler, note this is not the same as http.HandlerFunc
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// ResponseHeader is a middleware handler that adds a header to the response
type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

// ServeHTTP handles the request by adding the response header
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//add the header
	w.Header().Add(rh.headerName, rh.headerValue)
	//call the wrapped handler
	rh.handler.ServeHTTP(w, r)
}

// NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handlerToWrap http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}

// Simple example of a middleware response
type PrintHello struct {
	handler func(http.ResponseWriter, *http.Request)
}

func (mid *PrintHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	mid.handler(w, r)
}

func PrintHelloMiddleware(handlerToWrap func(http.ResponseWriter, *http.Request)) *PrintHello {
	return &PrintHello{handler: handlerToWrap}
}

// Request contexts were introduced in Go version 1.7, and they allow several advanced techniques, but the one we are concerned with here is the storage of request-scoped values. The request context gives us a spot to store and retrieve key/value pairs that stay with the http.Request object. Since a new instance of that object is created at the start of every request, anything we put into it will be particular to the current request.

func UploadServer() {
	// https://github.com/TutorialEdge
	os.Setenv("ADDR", "127.0.0.1:8080")
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.Handle("/time", PrintHelloMiddleware(currentTimeHandler))
	mux.HandleFunc("/upload", uploadHandler)

	//wrap entire mux with logger middleware
	wrappedMux := NewLogger(NewResponseHeader(mux, "X-My-Header", "my header value"))

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
