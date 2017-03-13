package bunq

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Error represents an error returned by the bunq API.
type Error struct {
	ErrorDescription           string
	ErrorDescriptionTranslated string
}

// Errors is an array of Error structs.
type Errors []Error

func (e Errors) Error() string {
	errs := make([]string, len(e))
	for i, err := range e {
		errs[i] = err.ErrorDescription
	}

	return strings.Join(errs, ", ")
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
		return fmt.Errorf("could not decode errors from json: %v", err)
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
