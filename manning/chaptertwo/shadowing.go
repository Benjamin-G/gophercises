package chaptertwo

import (
	"log"
	"net/http"
)

func ShadowRunner() {
	err := listing1()
	if err != nil {
		log.Fatal(err)
	}
}

func listing1() error {
	// Client will not be assigned, due to scoping
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Printf("%[1]T %[1]v", client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Printf("%[1]T %[1]v", client)
	}

	_ = client
	log.Printf("%[1]T %[1]v", client)
	return nil
}

func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	log.Printf("%[1]T %[1]v", client)
	_ = client
	return nil
}

func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

// BEST WAY if SHARED Error Handling
func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		// Common error handling
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	c := &http.Client{}
	return c, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
