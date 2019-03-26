package vonigo

import (
	"encoding/json"
	"errors"
)

type errorResponse struct {
	ErrNo  int           `json:"errNo"`
	ErrMsg string        `json:"errMsg"`
	Errors []vonigoError `json:"Errors"`
}

type vonigoError struct {
	FieldID   int    `json:"fieldID"`
	FieldName string `json:"fieldName"`
	ErrNo     int    `json:"errNo"`
	ErrMsg    string `json:"errMsg"`
}

// checkVonigoError Returns actual error codes generated by Vonigo API regardless of response code(Vonigo always returns 200 :-|)
func checkVonigoError(body []byte) error {
	errorResponse := errorResponse{}

	err := json.Unmarshal(body, &errorResponse)

	if err != nil {
		return err
	}

	if errorResponse.ErrMsg != "" || errorResponse.ErrNo != 0 {
		errString := "The following errors were returned by the API: "

		for _, errItem := range errorResponse.Errors {
			errString += errItem.ErrMsg
		}

		return errors.New(errString + " ")
	}

	return nil

}
