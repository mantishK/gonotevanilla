package validate

import (
	"errors"
)

//checks if the data contains the keys
func Required(data interface{}, keys []string) error {
	if data == nil {
		return errors.New("data not set")
	}
	dataMap := data.(map[string]interface{})
	for _, value := range keys {
		_, ok := dataMap[value]
		if !ok {
			return errors.New(value + " required ")
		}
	}
	return nil
}
