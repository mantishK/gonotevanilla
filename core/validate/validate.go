package validate

import (
	"errors"
	// "fmt"
	"net/url"
)

//checks if the data contains the keys. Used for POST, PUT and DELETE requests
func RequiredData(data map[string]interface{}, keys []string) (int, error) {
	if data == nil {
		return 0, errors.New("data not set")
	}
	for count, value := range keys {
		_, ok := data[value]
		if !ok {
			return count, errors.New(value + " required ")
		}
	}
	return -1, nil
}

func RequiredParams(data url.Values, keys []string) error {
	if data == nil {
		return errors.New("data not set")
	}
	for count, value := range keys {
		_, ok := data[value]
		if !ok {
			return count, errors.New(value + " required ")
		}
	}
	return nil
}
