package validate

import (
	"errors"
)

//checks if the obj contains the keys
func Required(obj map[string]string, keys []string) errors {
	for key := range keys {
		k, err := obj[key]
		if err {
			return errors.New(key + " does not exist ")
		}
	}
}
