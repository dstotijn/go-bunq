package bunq

import (
	"encoding/json"
	"fmt"
	"io"
)

// Error represents an error returned by the bunq API.
type Error struct {
	ErrorDescription           string
	ErrorDescriptionTranslated string
}

// Errors is an array of Error structs.
type Errors []Error

func (e Errors) Error() string {
	var errs []string
	for i := range e {
		errs = append(errs, e[i].ErrorDescription)
	}

	return fmt.Sprintf("%v", errs)
}

func decodeError(r io.Reader) error {
	var apiError struct {
		Error []struct {
			ErrorDescription           string `json:"error_description"`
			ErrorDescriptionTranslated string `json:"error_description_translated"`
		} `json:"Error"`
	}
	err := json.NewDecoder(r).Decode(&apiError)
	if err != nil {
		return fmt.Errorf("bunq: could not decode errors from json: %v", err)
	}

	var errors Errors
	for i := range apiError.Error {
		errors = append(errors, Error{
			ErrorDescription:           apiError.Error[i].ErrorDescription,
			ErrorDescriptionTranslated: apiError.Error[i].ErrorDescriptionTranslated,
		})
	}

	return errors
}
