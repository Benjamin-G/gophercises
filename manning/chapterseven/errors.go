package chapterseven

import (
	"fmt"
	"log"
)

func bar() error {
	return barError{}
}

type barError struct{}

func (b barError) Error() string {
	return "bar error <--"
}

func runError() error {
	err := bar()
	if err != nil {
		//  fmt.Printf("err: %T\n", err)
		//  return fmt.Errorf("bar failed: %v", err)
		//return fmt.Errorf("bar failed: %w", err)
		return fmt.Errorf("bar failed: %w", barError{})
	}

	// ...
	return nil
}

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func runError2() error {
	err := bar()
	if err != nil {
		fmt.Printf("err: %T\n", err)
		fmt.Printf("err: %T\n", transientError{err: err})

		return fmt.Errorf("failed to get transaction %w", transientError{err: err})
	}

	// ...
	return nil
}

type Route struct{}

func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate source coordinates: %w", err)
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate target coordinates: %w", err)
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}
