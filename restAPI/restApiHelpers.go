package restAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"strings"
)

// ********************************************************************************
// extractAndValidateJsonBody
func extractAndValidateJsonBody(responseWriterPointer *http.ResponseWriter, httpRequest *http.Request, myInputTypeVariable interface{}) (err error) {
	// If the Content-Type header is present, check that it has the value
	// application/json. Note that we are using the gddo/httputil/header
	// package to parse and extract the value here, so the check works
	// even if the client includes additional charset or boundary
	// information in the header.
	responseWriter := *responseWriterPointer
	if httpRequest.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(httpRequest.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(responseWriter, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	httpRequest.Body = http.MaxBytesReader(responseWriter, httpRequest.Body, 1048576)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(httpRequest.Body)
	dec.DisallowUnknownFields()

	var p = myInputTypeVariable //RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct
	err = dec.Decode(&p)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(responseWriter, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(responseWriter, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Println(err.Error())
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return err
	}

	// Call decode again, using a pointer to an empty anonymous struct as
	// the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		http.Error(responseWriter, msg, http.StatusBadRequest)
		return
	}

	//fmt.Fprintf(responseWriter, "RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct: %+v", p)

	return nil
}

// ********************************************************************************
// SetLogger
// Set to use the same logger reference as is used by central part of system
func (restAPI *RestApiStruct) SetLogger(logger *logrus.Logger) {
	restAPI.logger = logger

	return

}

// ********************************************************************************
// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (restAPI *RestApiStruct) SetDialAddressString(dialAddress string) {
	restAPI.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}
